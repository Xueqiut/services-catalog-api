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