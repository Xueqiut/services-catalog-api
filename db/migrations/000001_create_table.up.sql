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

INSERT INTO services 
  (name, description, user_id) 
VALUES 
  ('Locate Us', 'The location service', 1),
  ('Contact Us', 'The contact service', 1),
  ('Notifications', 'The notifications service', 1),
  ('Reporting', 'The reporting service', 1);

INSERT INTO versions
  (name, service_id, enabled) 
VALUES 
  ('v1', 1, true),
  ('v1', 2, false),
  ('v2', 2, true),
  ('v1', 3, false),
  ('v2', 3, false),
  ('v3', 3, true),
  ('v1', 4, false),
  ('v2', 4, false),
  ('v3', 4, false),
  ('v4', 4, true);