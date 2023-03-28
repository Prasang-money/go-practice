# Step to install MySql docker container.
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=money -p 3306:3306 -d mysql
docker exec -it container_Id bash
musql -u root -p
After this you will get a promt asking for password.
- Sql statement to create database name snippetbox
 CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
- Switch to using database snippetbox
 USE snippetbox;
- Sql statement for creating snippets table.
CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);
 Add an index on created coloumn.
 CREATE INDEX idx_snippets_created ON snippets(created);


- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY));
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
Creating a New User

- Run the below command to create a user name with web.
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
Once that’s done type exit to leave the MySQL prompt.
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– N
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

- command to create a new user with Select and Insert privilage.
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

