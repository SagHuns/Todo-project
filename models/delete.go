package models

import "github.com/SagHuns/todo-project/db"

func Delete(id int64) (int64, error){
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	// Fecha o DB qunado a operação encerrar
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
