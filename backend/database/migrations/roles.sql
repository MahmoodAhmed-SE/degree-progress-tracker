BEGIN;

INSERT INTO roles (domain, name) VALUES ('admin', 'USERS_READ');
INSERT INTO roles (domain, name) VALUES ('admin', 'USERS_CREATE');
INSERT INTO roles (domain, name) VALUES ('admin', 'USERS_UPDATE');
INSERT INTO roles (domain, name) VALUES ('admin', 'USERS_DELETE');

INSERT INTO roles_group (role_id, group_id) VALUES(
  ((SELECT id FROM roles WHERE name='USERS_READ'),1)
  ((SELECT id FROM roles WHERE name='USERS_CREATE'),1)
  ((SELECT id FROM roles WHERE name='USERS_UPDATE'),1)
  ((SELECT id FROM roles WHERE name='USERS_DELETE'),1)
  )

COMMIT;
