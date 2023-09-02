package models

import "github.com/SagHuns/todo-project/db"

func GetAll() (todos []Todo, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // tratar no handler, não aqui
	}
 
	// Fecha o DB qunado a operação encerrar
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue  // Pula para o próximo item do DB
		}

		todos = append(todos, todo)
	}

	return
}