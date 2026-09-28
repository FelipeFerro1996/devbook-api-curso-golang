insert into usuarios(nome, nick, email, senha)
values
("usuario 1", "usuario_1", "usuario_1@gmail.com", "$2a$10$qffVUAoWoSo/hQ9kqRepO.Vw/ADBKOfhXfRYgrvHJKliz/UfA6UM2"),
("usuario 2", "usuario_2", "usuario_2@gmail.com", "$2a$10$qffVUAoWoSo/hQ9kqRepO.Vw/ADBKOfhXfRYgrvHJKliz/UfA6UM2"),
("usuario 3", "usuario_3", "usuario_3@gmail.com", "$2a$10$qffVUAoWoSo/hQ9kqRepO.Vw/ADBKOfhXfRYgrvHJKliz/UfA6UM2");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);