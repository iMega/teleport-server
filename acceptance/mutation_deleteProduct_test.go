package acceptance

import (
	"encoding/json"

	"github.com/imega/teleport-server/acceptance/helper"
	"github.com/imega/teleport-server/api"
	"github.com/imega/teleport-server/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Тестирование удаления ноды, через удаление продукта", func() {
	var (
		ownerID = string(uuid.NewUUID())
		product api.Product
	)

	Context("Каталог пустой, создание товара", func() {
		It("товар создан", func() {
			product = helper.CreateProduct(ownerID)
		})
	})

	Context("Удаление товара", func() {
		It("товар удален", func() {
			vars := map[string]string{
				"id": product.GetId(),
			}
			body, bc := helper.RequestGraphQL(ownerID, "fixture/delete_product.graphql", vars, "DeleteProduct")
			defer bc()

			type (
				Data struct {
					Deleted bool `json:"deleteProduct"`
				}
				Response struct {
					Data Data `json:"data"`
				}
			)
			resp := &Response{}
			err := json.Unmarshal(body, resp)
			Expect(err).NotTo(HaveOccurred())

			Expect(resp.Data.Deleted).To(Equal(true))
		})
	})
})
