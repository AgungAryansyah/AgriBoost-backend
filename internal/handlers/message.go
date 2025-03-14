package handlers

import (
	"AgriBoost/internal/infra/middleware"
	ws "AgriBoost/internal/infra/websocket"
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageHandler struct {
	messageService   services.MessageServiceItf
	communityService services.CommunityServiceItf
	userService      services.UserServiceItf
	validator        *validator.Validate
	middleware       middleware.MiddlewareItf
}

func NewMessageHandler(routerGroup fiber.Router, messageService services.MessageServiceItf, communityService services.CommunityServiceItf, userService services.UserServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	MessageHandler := MessageHandler{
		messageService:   messageService,
		communityService: communityService,
		userService:      userService,
		validator:        validator,
		middleware:       middleware,
	}

	routerGroup = routerGroup.Group("/message")
	routerGroup.Get("/ws/:roomID", middleware.Authentication, websocket.New(MessageHandler.MessageWebSocketHandler))
	routerGroup.Post("/get", middleware.Authentication, MessageHandler.GetMessages)
}

func (m *MessageHandler) GetMessages(ctx *fiber.Ctx) error {
	var param dto.MessageParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := m.validator.Struct(param); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var messages []dto.MessageDto
	if err := m.messageService.GetMessages(&messages, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", messages)
}

func (m *MessageHandler) MessageWebSocketHandler(c *websocket.Conn) {
	roomID := c.Params("roomID")

	roomUUID, err := uuid.Parse(roomID)
	if err != nil {
		return
	}

	var exist bool
	if err := m.communityService.IsCommunityExist(&exist, roomUUID); err != nil || !exist {
		return
	}

	clientID := c.RemoteAddr().String()

	userID := c.Locals("userID").(string)
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return
	}

	var username string
	if err := m.userService.IsUserExistName(username, userUUID); err != nil {
		return
	}

	ws.RoomsMutex.Lock()
	if _, exists := ws.Rooms[roomID]; !exists {
		ws.Rooms[roomID] = &ws.Room{Clients: make(map[string]*ws.Client)}
	}
	room := ws.Rooms[roomID]
	ws.RoomsMutex.Unlock()

	room.Mutex.Lock()
	room.Clients[clientID] = &ws.Client{ID: clientID, Name: username, Conn: c, Room: roomID}
	room.Mutex.Unlock()

	for {
		var msg string
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Read error:", err)
			break
		}

		send := dto.SendMessage{
			Message:     msg,
			CommunityId: roomUUID,
			UserId:      userUUID,
		}

		if err := m.validator.Struct(send); err != nil {
			return
		}

		if err := m.messageService.SendMessage(send); err != nil {
			return
		}

		room.Mutex.Lock()
		for _, client := range room.Clients {
			if err := client.Conn.WriteJSON(fiber.Map{
				"user_id": userID,
				"sender":  username,
				"message": msg,
			}); err != nil {
				log.Println("Write error:", err)
			}
		}
		room.Mutex.Unlock()
	}

	room.Mutex.Lock()
	delete(room.Clients, clientID)
	room.Mutex.Unlock()
}
