package models

import "github.com/SagHuns/todo-project/db"

func Get(id int64) (todo Todo, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // tratar no handler, não aqui
	}
 
	// Fecha o DB qunado a operação encerrar
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}