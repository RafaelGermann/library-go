package repositories

import (
	"api/src/models"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func PublicacaoRepository(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

func (repositorio publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO Publicacoes (Titulo, Conteudo, Autor_Id) VALUES (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	result, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	idInserido, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(idInserido), nil
}

func (repositorio publicacoes) BuscarPorId(publicacaoId uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query("SELECT P.id, titulo, conteudo, P.autor_id, U.nick, curtidas, P.criadoEm FROM Publicacoes P INNER JOIN Usuarios U ON U.ID = autor_id WHERE P.id = ?", publicacaoId)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}
	return publicacao, nil
}
