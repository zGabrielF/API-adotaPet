package banco

import (
	"apiAdotaPet/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Close(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil

}
