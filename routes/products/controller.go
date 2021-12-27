package products

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MiguelBarriosC/goProducts/db"     // Package products
	"github.com/MiguelBarriosC/goProducts/models" // Pakage models
	guuid "github.com/google/uuid"                // Genera ids
)

var Table string = "products"

func GetAll() (us []models.Product, err error) {
	db := db.GetConnection()                        // Conexión a la base de datos.
	rows, err := db.Query("SELECT * FROM " + Table) // Query

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close() // Cierra el recurso

	productos := []models.Product{}

	/* El método Next retorna un bool, mientras sea true indicará que existe
	un valor siguiente para leer.*/
	for rows.Next() {
		p := new(models.Product)
		/* Escanea el valor actual de la fila e inserta el retorno
		en los campos que corresponden.*/
		err := rows.Scan(&p.Product_id, &p.Name, &p.Description, &p.Status, &p.Creation_date, &p.Update_date, &p.Account_id, &p.Format_product, &p.Value_unit, &p.Unit_name, &p.Unit_description, &p.Stock)
		if err != nil {
			fmt.Println(err)

		}
		productos = append(productos, *p)
	}
	return productos, nil
}

func GetOne(id string) (models.Product, error) {
	db := db.GetConnection() // Conexión a la base de datos.

	q := "SELECT * FROM " + Table + " WHERE product_id=?" // Query

	p := models.Product{}
	/* Escanea el valor actual de la fila e inserta el retorno
	en los campos que corresponden.*/
	err := db.QueryRow(q, id).Scan(&p.Product_id, &p.Name, &p.Description, &p.Status, &p.Creation_date, &p.Update_date, &p.Account_id, &p.Format_product, &p.Value_unit, &p.Unit_name, &p.Unit_description, &p.Stock)

	if err != nil {
		return models.Product{}, err
	}

	return p, nil
}

//Para insertar format_product, la estructura es de tipo interface "{\"data\": \"example\"}"
func Create(p models.Product) error {

	if len(strings.TrimSpace(p.Product_id)) == 0 {
		id := guuid.New()
		p.Product_id = id.String()
	}

	db := db.GetConnection() // Conexión a la base de datos.

	// Query para insertar los datos en la tabla products
	q := "INSERT INTO " + Table + " (product_id, name, description, status, creation_date, update_date, account_id, format_product, value_unit, unit_name, unit_description,stock) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"

	// Prepara la petición
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Ejecuta la petición
	r, err := stmt.Exec(p.Product_id, p.Name, p.Description, p.Status, time.Now(), time.Now(), p.Account_id, p.Format_product, p.Value_unit, p.Unit_name, p.Unit_description, p.Stock)
	if err != nil {
		return err
	}
	/*Confirma que una fila fuera afectada
	en caso contrario devuelve un error.*/
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	// Retorna un nil para confirmar que no existe un error.
	return nil
}

func Update(p models.Product) error {

	_, e := GetOne(p.Product_id)
	if e != nil {
		return errors.New("no existe")
	}

	db := db.GetConnection() // Conexión a la base de datos.

	// Query para insertar los datos en la tabla notes
	q := "UPDATE " + Table + " SET name=?, description=?, status=?, update_date=?, account_id=?, format_product=?, value_unit=?, unit_name=?, unit_description=?, stock=? WHERE product_id=?"

	// Prepara la petición
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Ejecuta la petición pasando los datos correspondientes.
	r, err := stmt.Exec(p.Name, p.Description, p.Status, time.Now(), p.Account_id, p.Format_product, p.Value_unit, p.Unit_name, p.Unit_description, p.Stock, p.Product_id)
	if err != nil {
		return err
	}
	/*Confirma que una fila fuera afectada
	en caso contrario devuelve un error.*/
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	// Retorna un nil para confirmar que no existe un error.
	return nil
}

func Delete(id string) bool {
	db := db.GetConnection() // Conexión a la base de datos.

	q := `DELETE FROM ` + Table + ` WHERE product_id=?`

	stmt, err := db.Prepare(q) // Prepara la petición
	if err != nil {
		return false
	}
	defer stmt.Close()

	r, err := stmt.Exec(id) // Ejecuta la petición
	if err != nil {
		return false
	}
	/*Confirma que una fila fuera afectada
	en caso contrario devuelve un false.*/
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return false
	}
	return true // Retorna true para inidicar que la operación se realizo correctamente.
}
