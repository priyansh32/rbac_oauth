-- Create a table for users
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);
-- Insert sample users
INSERT INTO users (username, email, password_hash)
VALUES (
        "harry_potter",
        "harry@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    ),
    (
        "hermione_granger",
        "hermione@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    ),
    (
        "ron_weasley",
        "ron@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    ),
    (
        "albus_dumbledore",
        "dumbledore@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    ),
    (
        "severus_snape",
        "snape@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    ),
    (
        "luna_lovegood",
        "luna@example.com",
        "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
    );
-- Create documents table
CREATE TABLE documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type VARCHAR(255),
    owner INTEGER,
    -- Assuming this is a foreign key referencing the users table
    title VARCHAR(255),
    content TEXT
);
-- Insert additional sample documents
INSERT INTO documents (type, owner, title, content)
VALUES (
        "public",
        "1",
        "The Forbidden Forest",
        "The forest was dark and deep, a place of mystery and danger."
    ),
    (
        "private",
        "2",
        "Advanced Potion-Making",
        "A detailed guide to advanced potion-making, annotated by Severus Snape."
    ),
    (
        "public",
        "3",
        "Quidditch Through the Ages",
        "The definitive guide to the history of Quidditch, from its origins to the modern game."
    ),
    (
        "protected",
        "4",
        "The Mirror of Erised",
        "The mirror shows the deepest, most desperate desire of our hearts."
    ),
    (
        "private",
        "5",
        "The Half-Blood Prince's Notes",
        "Personal notes and observations on potion-making, with a few extra comments."
    ),
    (
        "public",
        "6",
        "Thestrals and Nargles",
        "An exploration of magical creatures, invisible to those who have not witnessed death."
    ),
    (
        "protected",
        "1",
        "The Marauder\'s Map",
        "A magical map that reveals the layout of Hogwarts and the location of everyone inside it."
    ),
    (
        "public",
        "2",
        "The Library of Alexandria",
        "A historical account of the famous library, exploring its magical and non-magical aspects."
    ),
    (
        "private",
        "3",
        "A Brief History of Quidditch Teams",
        "An insider\'s look at the rise and fall of famous Quidditch teams throughout history."
    ),
    (
        "protected",
        "4",
        "The Elder Wand",
        "The tale of the Deathly Hallows and the most powerful wand in existence."
    ),
    (
        "public",
        "5",
        "The Art of Occlumency",
        "Mastering the skill of shielding one\'s mind from magical intrusion."
    ),
    (
        "private",
        "6",
        "Fantastic Beasts and Where to Find Them",
        "An in-depth study of magical creatures, from Acromantulas to the elusive Niffler."
    ),
    (
        "public",
        "1",
        "Defense Against the Dark Arts Handbook",
        "Essential spells and techniques for defending against dark creatures and curses."
    ),
    (
        "protected",
        "2",
        "Time-Turner Chronicles",
        "The secrets and limitations of the Time-Turner, a magical device capable of time travel."
    ),
    (
        "public",
        "3",
        "The Great Wizarding Wars",
        "A historical account of the conflicts and battles that shaped the wizarding world."
    ),
    (
        "private",
        "4",
        "The Philosopher\;s Stone",
        "The quest for the Philosopher\'s Stone and its potential to grant immortality."
    ),
    (
        "protected",
        "5",
        "Potions Master\'s Journal",
        "A collection of advanced potion recipes and Snape\'s personal notes on potion-making."
    ),
    (
        "public",
        "6",
        "The Quibbler\'s Conspiracy Theories",
        "Unconventional theories and insights into magical mysteries, as published in The Quibbler."
    ),
    (
        "private",
        "1",
        "The Invisibility Cloak",
        "The legend and history of the Invisibility Cloak, a powerful magical garment."
    ),
    (
        "protected",
        "2",
        "The Room of Requirement",
        "A guide to the magical room that transforms itself to suit the needs of the seeker."
    );
CREATE TABLE clients (
    id varchar(64) PRIMARY KEY,
    secret VARCHAR(64) NOT NULL,
    role VARCHAR(50) NOT NULL,
    redirect_uri text NOT NULL
);
INSERT INTO clients (id, secret, role, redirect_uri)
VALUES (
        "laama",
        "secretkeythis",
        "editor",
        "localhost:3000/auth/callback"
    );
CREATE TABLE authorization_codes (
    code VARCHAR(64) PRIMARY KEY,
    client_id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    code_challenge VARCHAR(64) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);