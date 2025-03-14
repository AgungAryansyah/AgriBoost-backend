CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profile_picture TEXT UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(15) NOT NULL, UNIQUE,
    password VARCHAR(255) NOT NULL,
    quiz_point INTEGER NOT NULL DEFAULT 0,
    donation_point INTEGER NOT NULL DEFAULT 0,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS communities (
    community_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS quizzes (
    quiz_id UUID PRIMARY KEY,
    theme VARCHAR(50) NOT NULL,
    title VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS campaigns (
    campaign_id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    goal_amount INTEGER NOT NULL CHECK (goal_amount > 0), 
    collected_amount INTEGER NOT NULL DEFAULT 0, 
    is_active BOOLEAN NOT NULL DEFAULT TRUE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS donations (
    donation_id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL,
    user_id UUID NOT NULL,
    amount INTEGER NOT NULL CHECK (amount > 0), 
    status VARCHAR(20) NOT NULL DEFAULT 'pending', 
    transaction_id VARCHAR(100) NOT NULL UNIQUE,
    donation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(campaign_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS questions (
    question_id UUID PRIMARY KEY, 
    quiz_id UUID NOT NULL,
    score INTEGER NOT NULL CHECK (score > 0), 
    question_text TEXT NOT NULL,
    question_image TEXT,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(quiz_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS question_options (
    option_id UUID PRIMARY KEY,
    question_id UUID NOT NULL, 
    is_correct BOOLEAN NOT NULL DEFAULT FALSE,
    option_text TEXT NOT NULL,
    option_image TEXT,
    FOREIGN KEY (question_id) REFERENCES questions(question_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS quiz_attempts (
    attempt_id UUID PRIMARY KEY, 
    user_id UUID NOT NULL,
    quiz_id UUID NOT NULL,
    total_score INTEGER NOT NULL CHECK (total_score > 0),
    finished_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(quiz_id) ON DELETE CASCADE
);



CREATE TABLE IF NOT EXISTS community_members (
    member_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    community_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (community_id) REFERENCES communities(community_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS articles (
    article_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(50) NOT NULL,
    content_text TEXT NOT NULL,
    content_image TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users (id, name, email, password, quiz_point, donation_point, is_admin) VALUES
    (gen_random_uuid(), 'Admin User', 'admin@example.com', crypt('adminpassword', gen_salt('bf')), 100, 50, TRUE),
    (gen_random_uuid(), 'John Doe', 'john@example.com', crypt('johnspassword', gen_salt('bf')), 50, 20, FALSE),
    (gen_random_uuid(), 'Jane Smith', 'jane@example.com', crypt('janespassword', gen_salt('bf')), 70, 30, FALSE);

INSERT INTO communities (community_id, name, description) VALUES
    (gen_random_uuid(), 'Farmers United', 'A community for farmers to collaborate and share knowledge'),
    (gen_random_uuid(), 'Organic Growers', 'A space for organic farming discussions');

INSERT INTO quizzes (quiz_id, theme, title) VALUES
    (gen_random_uuid(), 'Agriculture', 'Farm Knowledge Test'),
    (gen_random_uuid(), 'Sustainability', 'Eco-Friendly Practices');

INSERT INTO campaigns (campaign_id, title, description, goal_amount, user_id) VALUES
    (gen_random_uuid(), 'Help Small Farmers', 'Support local farmers with resources', 5000, (SELECT id FROM users WHERE email='admin@example.com')),
    (gen_random_uuid(), 'Organic Farming Fund', 'Fund sustainable organic farming projects', 7000, (SELECT id FROM users WHERE email='john@example.com'));

INSERT INTO donations (donation_id, campaign_id, user_id, amount, status, transaction_id) VALUES
    (gen_random_uuid(), (SELECT campaign_id FROM campaigns LIMIT 1), (SELECT id FROM users WHERE email='jane@example.com'), 200, 'completed', 'txn12345'),
    (gen_random_uuid(), (SELECT campaign_id FROM campaigns LIMIT 1 OFFSET 1), (SELECT id FROM users WHERE email='john@example.com'), 500, 'pending', 'txn67890');

INSERT INTO questions (question_id, quiz_id, score, question_text) VALUES
    (gen_random_uuid(), (SELECT quiz_id FROM quizzes LIMIT 1), 10, 'What is the most commonly grown crop worldwide?'),
    (gen_random_uuid(), (SELECT quiz_id FROM quizzes LIMIT 1 OFFSET 1), 15, 'Which method is best for soil conservation?');

INSERT INTO question_options (option_id, question_id, is_correct, option_text) VALUES
    (gen_random_uuid(), (SELECT question_id FROM questions LIMIT 1), TRUE, 'Wheat'),
    (gen_random_uuid(), (SELECT question_id FROM questions LIMIT 1), FALSE, 'Corn'),
    (gen_random_uuid(), (SELECT question_id FROM questions LIMIT 1 OFFSET 1), TRUE, 'Crop rotation'),
    (gen_random_uuid(), (SELECT question_id FROM questions LIMIT 1 OFFSET 1), FALSE, 'Monocropping');

INSERT INTO quiz_attempts (attempt_id, user_id, quiz_id, total_score) VALUES
    (gen_random_uuid(), (SELECT id FROM users WHERE email='john@example.com'), (SELECT quiz_id FROM quizzes LIMIT 1), 10),
    (gen_random_uuid(), (SELECT id FROM users WHERE email='jane@example.com'), (SELECT quiz_id FROM quizzes LIMIT 1 OFFSET 1), 15);

INSERT INTO community_members (member_id, user_id, community_id) VALUES
    (gen_random_uuid(), (SELECT id FROM users WHERE email='john@example.com'), (SELECT community_id FROM communities LIMIT 1)),
    (gen_random_uuid(), (SELECT id FROM users WHERE email='jane@example.com'), (SELECT community_id FROM communities LIMIT 1 OFFSET 1));

INSERT INTO articles (article_id, user_id, title, content_text, content_image) VALUES
    (gen_random_uuid(), (SELECT id FROM users WHERE email='admin@example.com'), 'Sustainable Farming', 'A guide to sustainable farming techniques.', 'image1.jpg'),
    (gen_random_uuid(), (SELECT id FROM users WHERE email='jane@example.com'), 'Organic Certification', 'How to get certified as an organic farmer.', 'image2.jpg');
