package acceptance

import (
	"fmt"

	"github.com/imega/teleport-server/acceptance/helper"
	"github.com/imega/teleport-server/uuid"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Тестирование создания ноды, через создание продукта", func() {
	var (
		ownerID = string(uuid.NewUUID())
	)
	Context("Каталог пустой", func() {
		It("", func() {
			body, bc := helper.RequestGraphQL(ownerID, "fixture/create_product.graphql", nil, "CreateProduct")
			defer bc()

			fmt.Println(string(body))
		})
	})
})
