package products

import (
	"fmt"
	"testing"

	"github.com/MiguelBarriosC/goProducts/db"
	"github.com/MiguelBarriosC/goProducts/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Products Suite")
}

var _ = BeforeSuite(func() {
	db.Datab = "../../products_test.db"

	u := models.Product{
		Product_id:       "test_delete",
		Name:             "delete product",
		Description:      "description text",
		Status:           "good",
		Account_id:       "account_id test",
		Format_product:   "{\"name\":\"test_product\"}",
		Value_unit:       20.55,
		Unit_name:        "pechuga",
		Unit_description: "cuello",
		Stock:            34,
	}

	err := Create(u)

	if err != nil {
		fmt.Print("[ERROR] ", err)
	}
})

var _ = Describe("Productos", func() {

	Context("Insertar producto", func() {
		//Para insertar products.format_product, la estructura es de tipo interface "{\"data\": \"example\"}"
		It("Insertar producto completo sin id", func() {
			u := models.Product{
				Name:             "insert product",
				Description:      "description insert text",
				Status:           "good",
				Account_id:       "account_id insert test",
				Format_product:   "{\"op\":\"test_product insert\"}",
				Value_unit:       1.55,
				Unit_name:        "producto prueba",
				Unit_description: "producto de prueba",
				Stock:            14,
			}

			err := Create(u)
			Expect(err).Should(BeNil())
		})
	})

	Context("Listado de productos", func() {
		It("Obtener listado de productos", func() {
			us, err := GetAll()
			//La cantidad de elementos en us no debe ser 0.
			Expect(len(us)).ShouldNot(Equal(0))
			//No debe existir error
			Expect(err).Should(BeNil())
		})
	})

	Context("Obtener un producto", func() {
		It("Obtener un producto que existe", func() {

			var id string = "test_getOne"
			p, err := GetOne(id)

			expected_product := models.Product{
				Product_id:       "test_getOne",
				Name:             "pescado",
				Description:      "pescado argentino",
				Status:           "good",
				Creation_date:    "0001-01-01T00:00:00Z",
				Update_date:      "0001-01-01T00:00:00Z",
				Account_id:       "acountid",
				Format_product:   "{\"name\":\"pollito\"}",
				Value_unit:       100.33,
				Unit_name:        "lomo",
				Unit_description: "escama",
				Stock:            19,
			}
			Expect(err).Should(BeNil())
			Expect(p).Should(Equal(expected_product))
		})
		It("Obtener un producto que no existe", func() {
			id := "noexiste"

			us, err := GetOne(id)

			expected_user := models.Product{
				Product_id:       "",
				Name:             "",
				Description:      "",
				Status:           "",
				Creation_date:    "",
				Update_date:      "",
				Account_id:       "",
				Format_product:   nil,
				Value_unit:       0,
				Unit_name:        "",
				Unit_description: "",
				Stock:            0,
			}

			Expect(us).Should(Equal(expected_user))
			Expect(err.Error()).Should(ContainSubstring("no rows in result set"))
		})
	})

	Context("Editar producto", func() {
		It("Editar un producto que no existe", func() {
			u := models.Product{
				Product_id:       "test_update_fail",
				Name:             "test updated",
				Description:      "description update text",
				Status:           "good",
				Account_id:       "account_id update test",
				Format_product:   "{\"update\":\"test_product update\"}",
				Value_unit:       9.99,
				Unit_name:        "update product test",
				Unit_description: "test update",
				Stock:            999,
			}

			err := Update(u)
			//fmt.Print("ERROR ", err.Error())
			Expect(err.Error()).Should(ContainSubstring("no existe"))
		})
		It("Editar un producto que existe", func() {
			p := models.Product{
				Product_id:       "test_update",
				Name:             "test updated",
				Description:      "description update text",
				Status:           "good",
				Account_id:       "account_id update test",
				Format_product:   "{\"update\":\"test_product update\"}",
				Value_unit:       9.99,
				Unit_name:        "update product test",
				Unit_description: "test update",
				Stock:            999,
			}

			err := Update(p)
			//No debera existir error en la operacion
			Expect(err).Should(BeNil())
		})
	})

	Context("Eliminar producto", func() {
		It("Eliminar un producto que no existe", func() {
			id := "noexiste"
			err := Delete(id)
			//Retorna false por que el usuario no existe.
			Expect(err).Should(BeFalse())
		})
		It("Eliminar un producto que existe", func() {
			// Elimina el producto creado en before_suite
			id := "test_delete"
			err := Delete(id)
			Expect(err).Should(BeTrue())
		})
	})
})
