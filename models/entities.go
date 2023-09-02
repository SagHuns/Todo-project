package models

// Struct para mandar para o DB como também para os requests
type Todo struct {
	ID 			int64 	`json:"id"`  // Significa que quando estiver no meu json ele será "id" minúsculo
	Title 		string 	`json:"title"`
	Description string 	`json:"description"`
	Done 		bool 	`json:"done"`
}