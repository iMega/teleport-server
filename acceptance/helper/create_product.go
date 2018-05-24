package helper

import (
	"encoding/json"

	"github.com/imega/teleport-server/api"
	. "github.com/onsi/gomega"
)

func CreateProduct(ownerID string) api.Product {
	body, bc := RequestGraphQL(ownerID, "fixture/create_product.graphql", nil, "CreateProduct")
	defer bc()

	type (
		Data struct {
			Product api.Product `json:"createProduct"`
		}
		Response struct {
			Data Data `json:"data"`
		}
	)
	resp := Response{}
	err := json.Unmarshal(body, &resp)
	Expect(err).NotTo(HaveOccurred())
	Expect(resp.Data.Product.Id).NotTo(Equal(""))

	return resp.Data.Product
}
