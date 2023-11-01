CREATE TABLE Settings (
  codify BOOLEAN DEFAULT false NOT NULL,
  model TEXT DEFAULT 'gpt-4' NOT NULL,
  editor TEXT DEFAULT 'nvim' NOT NULL,
  temp REAL DEFAULT 10 NOT NULL
);
CREATE TABLE Convo (
  id INTEGER PRIMARY KEY,
  title TEXT,
  slug TEXT,
  system TEXT
);
CREATE INDEX idx_convo_id ON Convo(id);
CREATE TABLE Messages (
  id INTEGER PRIMARY KEY,
  role TEXT NOT NULL,
  msg TEXT NOT NULL,
  convo_id INTEGER,
  FOREIGN KEY (convo_id) REFERENCES Convo(id)
);
CREATE INDEX idx_messages_convo_id ON Messages(convo_id);
CREATE TABLE Pins (
  id INTEGER PRIMARY KEY,
  convo_id INTEGER,
  FOREIGN KEY (convo_id) REFERENCES Convo(id)
);
CREATE INDEX idx_pins_convo_id ON Pins(convo_id);