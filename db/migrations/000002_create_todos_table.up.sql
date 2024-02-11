

BEGIN;
  CREATE TABLE IF NOT EXISTS todos(
    id VARCHAR (255) UNIQUE NOT NULL PRIMARY KEY,
    text TEXT,
    done BOOL DEFAULT FALSE,
    user_id VARCHAR (255)
  );
  CREATE INDEX on todos(id);
COMMIT;