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

insert into publicacoes(titulo,conteudo,autor_id)
values
("Publicação usuário 1", "Essa é uma publicação do usuario 1", 1),
("Publicação usuário 2", "Essa é uma publicação do usuario 2", 2),
("Publicação usuário 3", "Essa é uma publicação do usuario 3", 3);