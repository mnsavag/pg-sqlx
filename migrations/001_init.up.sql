CREATE TABLE IF NOT EXISTS decks 
(
    id TEXT UNIQUE,
    topic TEXT,
    description TEXT,
    links TEXT
);

CREATE TABLE IF NOT EXISTS cards
(
    id TEXT UNIQUE,
    question TEXT,
    answer TEXT
);

CREATE TABLE IF NOT EXISTS deck_cards
(
    id INTEGER PRIMARY KEY,
    card_id INTEGER,
    deck_id INTEGER,

    FOREIGN KEY (card_id) REFERENCES cards (id) ON DELETE CASCADE,
    FOREIGN KEY (deck_id) REFERENCES decks (id) ON DELETE CASCADE
);