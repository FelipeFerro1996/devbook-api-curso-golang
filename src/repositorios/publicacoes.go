package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositoriosDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) InserirPublicacao(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes(titulo, conteudo, autor_id) values(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacoes) BuscarPublicacoes(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select 
			p.titulo,
			p.conteudo,
			p.autor_id,
			u.nick,
			p.criadaem
		from
			publicacoes as p
		inner join 
			usuarios as u on p.autor_id = u.id
		where 
			p.autor_id = ?
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicaco modelos.Publicacao
		if erro := linhas.Scan(
			&publicaco.Titulo,
			&publicaco.Conteudo,
			&publicaco.AutorID,
			&publicaco.AutorNick,
			&publicaco.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicaco)
	}

	if erro := linhas.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select 
			p.*,
			u.nick
		from
			publicacoes as p
		inner join 
			usuarios as u on p.autor_id = u.id
		where
			p.id = ?
	`, publicacaoId)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicaco modelos.Publicacao
	if linha.Next() {
		if erro := linha.Scan(
			&publicaco.ID,
			&publicaco.Titulo,
			&publicaco.Conteudo,
			&publicaco.AutorID,
			&publicaco.Curtidas,
			&publicaco.CriadaEm,
			&publicaco.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicaco, nil
}
