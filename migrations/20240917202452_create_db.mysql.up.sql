CREATE TABLE `users`
(
    `id`        BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `email`     VARCHAR(255) NOT NULL UNIQUE,
    `name`      VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NULL,
    `password`  VARCHAR(255) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE workspaces
(
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255)    NOT NULL,
    owner_id   BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP       NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE projects
(
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name         VARCHAR(255)    NOT NULL,
    workspace_id BIGINT UNSIGNED NOT NULL,
    owner_id     BIGINT UNSIGNED NOT NULL,
    created_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL,
    FOREIGN KEY (workspace_id) REFERENCES workspaces (id),
    FOREIGN KEY (owner_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE boards
(
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name         VARCHAR(255)    NOT NULL,
    workspace_id BIGINT UNSIGNED NOT NULL,
    owner_id     BIGINT UNSIGNED NOT NULL,
    project_id   BIGINT UNSIGNED NULL,
    created_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL,
    FOREIGN KEY (workspace_id) REFERENCES workspaces (id),
    FOREIGN KEY (owner_id) REFERENCES users (id),
    FOREIGN KEY (project_id) REFERENCES projects (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE board_sections
(
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255)    NOT NULL,
    board_id   BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP       NULL,
    FOREIGN KEY (board_id) REFERENCES boards (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
