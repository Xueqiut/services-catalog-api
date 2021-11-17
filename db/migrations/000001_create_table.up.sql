CREATE TABLE IF NOT EXISTS services (
  id          serial PRIMARY KEY,
  name        VARCHAR(128) NOT NULL,
  description VARCHAR(255) NOT NULL,
  user_id     BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS versions (
  id          serial PRIMARY KEY,
  name        VARCHAR(128) NOT NULL,
  service_id  BIGINT,
  enabled     BOOLEAN,
  CONSTRAINT fk_service
      FOREIGN KEY(service_id) 
      REFERENCES services(id)
      ON DELETE CASCADE
);