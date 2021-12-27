package products

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MiguelBarriosC/goProducts/db"
	"github.com/MiguelBarriosC/goProducts/models"
	guuid "github.com/google/uuid"
)

var Table string = "products"

func GetAll() (us []models.Product, err error) {
	db := db.GetConnection()                        // Realizamos la conexión a la base de datos.
	rows, err := db.Query("SELECT * FROM " + Table) //Query

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	productos := []models.Product{}

	for rows.Next() {
		p := new(models.Product)

		err := rows.Scan(&p.Product_id, &p.Name, &p.Description, &p.Status, &p.Creation_date, &p.Update_date, &p.Account_id, &p.Format_product, &p.Value_unit, &p.Unit_name, &p.Unit_description, &p.Stock)
		if err != nil {
			fmt.Println(err)

		}
		productos = append(productos, *p)
	}
	return productos, nil
}

func GetOne(id string) (models.Product, error) {
	db := db.GetConnection() // Realizamos la conexión a la base de datos.

	q := "SELECT * FROM " + Table + " WHERE product_id=?" // Query

	p := models.Product{}
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

	db := db.GetConnection() // Realizamos la conexión a la base de datos.

	// Query para insertar los datos en la tabla products
	q := "INSERT INTO " + Table + " (product_id, name, description, status, creation_date, update_date, account_id, format_product, value_unit, unit_name, unit_description,stock) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"

	// Preparamos la petición
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Ejecutamos la petición
	r, err := stmt.Exec(p.Product_id, p.Name, p.Description, p.Status, time.Now(), time.Now(), p.Account_id, p.Format_product, p.Value_unit, p.Unit_name, p.Unit_description, p.Stock)
	if err != nil {
		return err
	}
	// Confirmamos que una fila fuera afectada
	// En caso contrario devolvemos un error.
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	// Retornamos un nil para confirmar que no existe un error.
	return nil
}

func Update(p models.Product) error {

	_, e := GetOne(p.Product_id)
	if e != nil {
		return errors.New("no existe")
	}

	db := db.GetConnection() // Realizamos la conexión a la base de datos.

	// Query para insertar los datos en la tabla notes
	q := "UPDATE " + Table + " SET name=?, description=?, status=?, update_date=?, account_id=?, format_product=?, value_unit=?, unit_name=?, unit_description=?, stock=? WHERE product_id=?"

	// Preparamos la petición
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Ejecutamos la petición pasando los datos correspondientes.
	r, err := stmt.Exec(p.Name, p.Description, p.Status, time.Now(), p.Account_id, p.Format_product, p.Value_unit, p.Unit_name, p.Unit_description, p.Stock, p.Product_id)
	if err != nil {
		return err
	}
	// Confirmamos que una fila fuera afectada
	// En caso contrario devolvemos un error.
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	// Retornamos un nil para confirmar que no existe un error.
	return nil
}

func Delete(id string) bool {
	db := db.GetConnection() // Realizamos la conexión a la base de datos.

	q := `DELETE FROM ` + Table + ` WHERE product_id=?`

	stmt, err := db.Prepare(q) // Preparamos la petición
	if err != nil {
		return false
	}
	defer stmt.Close()

	r, err := stmt.Exec(id) // Ejecutamos la petición
	if err != nil {
		return false
	}
	// Confirmamos que una fila fuera afectada
	// En caso contrario devolvemos un false.
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return false
	}
	return true // Retornamos true para inidicar que la operación se realizo correctamente.
}
