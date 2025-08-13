CREATE DATABASE IF NOT EXISTS tasking;
USE tasking;

DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS categories;

CREATE TABLE categories (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE tasks (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NULL,
  category_id INT NOT NULL,
  parent_task_id INT NULL,
  CONSTRAINT fk_tasks_category FOREIGN KEY (category_id) REFERENCES categories(id),
  CONSTRAINT fk_tasks_parent FOREIGN KEY (parent_task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

CREATE INDEX idx_parent ON tasks(parent_task_id);

INSERT INTO categories (name) VALUES ('Work'), ('Home');
INSERT INTO tasks (name, description, category_id, parent_task_id) VALUES
('Root task A', 'top-level', 1, NULL),
('Sub A1', 'child', 1, 1),
('Sub A2', 'child', 1, 1),
('Root task B', 'top-level', 2, NULL),
('Sub B1', 'child', 2, 4);
