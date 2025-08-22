ALTER TABLE projects
ADD COLUMN user_id BIGINT;

ALTER TABLE projects
ADD CONSTRAINT fk_projects_user
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE SET NULL;

CREATE INDEX idx_projects_user_id ON projects(user_id);
