CREATE TABLE IF NOT EXISTS decks 
(
    id TEXT UNIQUE NOT NULL,
    topic TEXT,
    description TEXT,
    links TEXT
);

CREATE TABLE IF NOT EXISTS cards
(
    id TEXT UNIQUE NOT NULL,
    question TEXT,
    answer TEXT
);

CREATE TABLE IF NOT EXISTS deck_cards
(
    id INTEGER PRIMARY KEY,
    card_id TEXT REFERENCES cards (id) ON DELETE CASCADE NOT NULL,
    deck_id TEXT REFERENCES decks (id) ON DELETE CASCADE NOT NULL
);