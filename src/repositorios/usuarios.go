package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositoriosDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // "%nomeOuNick%"

	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return usuarios, nil

}

func (repositorio Usuarios) BuscarPorID(usuarioId uint64) (modelos.Usuario, error) {

	var usuario modelos.Usuario
	linha, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioId)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	if linha.Next() {
		if erro := linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID)
	if erro != nil {
		return erro
	}

	return nil

}

func (repositorio Usuarios) Deletar(usuarioID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores(usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Usuarios) PararDeSeguir(usuario_id, seguidor_id uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario_id, seguidor_id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarTodosOsSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {

	// linhas, erro := repositorio.db.Query(
	// 	`select
	// 		nome,
	// 		email,
	// 		nick
	// 	from
	// 		usuarios
	// 	where
	// 		id in (
	// 			select
	// 				seguidor_id
	// 			from
	// 				seguidores
	// 			where
	// 				usuario_id = ?
	// 		)`,
	// 	usuarioId)
	linhas, erro := repositorio.db.Query(
		`select
			u.id, 
			u.nome, 
			u.email, 
			u.nick,
			u.criadoEm 
		from 
			usuarios as u
		inner join
			seguidores as s on u.id = s.seguidor_id where usuario_id = ?`,
		usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Nick,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return usuarios, nil

}

func (repositorio Usuarios) BuscarTodosOsSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		`select
			u.id, 
			u.nome, 
			u.email, 
			u.nick,
			u.criadoEm 
		from 
			usuarios as u
		inner join
			seguidores as s on u.id = s.usuario_id where seguidor_id = ?`,
		usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Nick,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return usuarios, nil

}

func (repositorio Usuarios) BuscarSenhaSalvaNoBanco(UsuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", UsuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(UsuarioId uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senha, UsuarioId); erro != nil {
		return erro
	}

	return nil
}
