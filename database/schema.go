package database


// CreateProductTable ...
func CreateProductTable() {

	DB.Query(`CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    amount integer,
    name text UNIQUE,
    description text,
    category text NOT NULL
)
`)

}