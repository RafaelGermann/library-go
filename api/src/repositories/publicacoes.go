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
	linha, erro := repositorio.db.Query("SELECT P.id, P.titulo, P.conteudo, P.autor_id, U.nick, P.curtidas, P.criadoEm FROM Publicacoes P INNER JOIN Usuarios U ON U.ID = autor_id WHERE P.id = ?", publicacaoId)
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

func (repositorio publicacoes) BuscarPublicacoes(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query("SELECT DISTINCT P.id, P.titulo, P.conteudo, P.autor_id, U.nick, P.Curtidas, P.criadoEm FROM Publicacoes P INNER JOIN Usuarios U on U.Id = P.autor_id INNER JOIN Seguidores S ON S.usuario_id = P.autor_id   WHERE U.Id = ? OR S.seguidor_id = ? ORDER BY 1 DESC", usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil

}

func (repositorio publicacoes) Atualizar(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE Publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio publicacoes) Deletar(id uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM Publicacoes WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio publicacoes) BuscarPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query("SELECT P.id, P.titulo, P.conteudo, P.autor_id, U.nick, P.Curtidas, P.criadoEm FROM Publicacoes P INNER JOIN Usuarios U ON U.Id = P.autor_id WHERE P.autor_Id = ? ORDER BY 1 DESC", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil

}

func (repositorio publicacoes) Curtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE Publicacoes SET curtidas = curtidas + 1  WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
func (repositorio publicacoes) Descurtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE Publicacoes SET curtidas = GREATEST(curtidas - 1, 0)  WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
