package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	. "github.com/onsi/gomega"
)

const defaultStyle = "\x1b[0m"
const cyanColor = "\x1b[36m"
const yellowColor = "\x1b[33m"

var (
	APIURL = "http://app:80"

	dumpReq = func(req *http.Request) {
		dump, err := httputil.DumpRequestOut(req, true)
		Expect(err).To(BeNil())
		fmt.Printf("%s\nREQUEST:\n%s\n%s\n", cyanColor, string(dump), defaultStyle)
	}

	dumpRes = func(res *http.Response) {
		dump, err := httputil.DumpResponse(res, true)
		Expect(err).To(BeNil())
		fmt.Printf("%s\nRESPONSE:\n%s\n%s\n", yellowColor, string(dump), defaultStyle)
	}
)

func RequestGraphQL(ownerID, queryFilename string, variables interface{}, operation string) ([]byte, func()) {
	//	dumper, err := config.GetConfigValue("DUMP_HTTP")
	//	Expect(err).NotTo(HaveOccurred())
	dumper := "1"

	b, err := ioutil.ReadFile(queryFilename)
	Expect(err).NotTo(HaveOccurred())

	ret, err := json.Marshal(struct {
		Query         string      `json:"query"`
		Variables     interface{} `json:"variables,omitempty"`
		OperationName string      `json:"operationName,omitempty"`
	}{
		Query:         string(b),
		Variables:     variables,
		OperationName: operation,
	})
	Expect(err).NotTo(HaveOccurred())

	req, err := http.NewRequest(http.MethodPost, APIURL, bytes.NewBuffer(ret))
	req.Header.Set("GRPC-METADATA-X-OWNER-ID", ownerID)

	if dumper != "" {
		dumpReq(req)
	}

	res, err := http.DefaultClient.Do(req)
	Expect(err).NotTo(HaveOccurred())
	Expect(res.StatusCode).To(Equal(http.StatusOK))

	if dumper != "" {
		dumpRes(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	Expect(err).NotTo(HaveOccurred())
	//Expect(body).NotTo(HaveGraphqlError())

	var bodyClose = func() {
		res.Body.Close()
	}

	return body, bodyClose
}
