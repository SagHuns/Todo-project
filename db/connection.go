// Arquivo responsável pela conexão com o banco de dados
package db

import (
	"example/hello/hello/configs"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

// Retorna o ponteiro para o DB e o erro 
func OpenConnection() (*sql.DB, error){
	conf := configs.GetDB()

	// sc utiliza Sprintf que faz com que cada %s seja substituída pelo seu conseguinte como conf.Host...
	// A ideia é criar uma string assim:
	// host=localhost port=5432 user=test password=test dbname=test sslmode=disabled
	// Para poder então realizar o sql.Open com o postgres e as configurações.
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
	conf.Host, conf.Port, conf.User, conf.Pass, conf.Database) 

	// Abrindo a conexão com o banco de dados - nesse caso é o postgres.
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	// Testa se a conexão foi sucedida
	err = conn.Ping()

	return conn, err
}
