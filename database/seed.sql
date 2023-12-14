DROP TABLE IF EXISTS played_game;

DROP TABLE IF EXISTS played_game_players;

DROP TABLE IF EXISTS game;

DROP TABLE IF EXISTS player;

CREATE TABLE game (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT
);

CREATE TABLE player (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT
);

CREATE TABLE played_game (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  date DATETIME,
  game_id INTEGER,
  winner_id INTEGER,
  FOREIGN KEY (game_id) REFERENCES game(id),
  FOREIGN KEY (winner_id) REFERENCES player(id)
);

CREATE TABLE played_game_players (
  played_game_id INTEGER,
  player_id INTEGER,
  PRIMARY KEY (played_game_id, player_id),
  FOREIGN KEY (played_game_id) REFERENCES played_game(id),
  FOREIGN KEY (player_id) REFERENCES player(id)
);

INSERT INTO game (name)
VALUES ('Nemesis'),
  ('Kingdom Death'),
  ('Wingspan');

INSERT INTO player (name)
VALUES ('Jimmy'),
  ('Laura'),
  ('Ylish');

INSERT INTO played_game (date, game_id, winner_id)
VALUES ('2023-12-13 12:00:00', 1, 2),
  ('2023-12-14 15:30:00', 2, 3),
  ('2023-12-15 18:45:00', 3, 1);

INSERT INTO played_game_players (played_game_id, player_id)
VALUES (1, 3),
  (1, 2),
  (1, 1);

INSERT INTO played_game_players (played_game_id, player_id)
VALUES (2, 3),
  (2, 1);

INSERT INTO played_game_players (played_game_id, player_id)
VALUES (3, 1),
  (3, 2);