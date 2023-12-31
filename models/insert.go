package models

import(
	"github.com/SagHuns/todo-project/db"
)

func Insert(todo Todo) (id int64, err error) {
	// Primeiro passo é tentar abrir uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return // tratar no handler, não aqui
	}
 
	// Fecha o DB qunado a operação encerrar
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}