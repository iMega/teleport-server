package acceptance_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var _ = BeforeSuite(func() {
	err := WaitForSystemUnderTestReady()
	Expect(err).NotTo(HaveOccurred())
})

func WaitForSystemUnderTestReady() error {
	cc, err := grpc.Dial("app:9000", grpc.WithInsecure())
	if err != nil {
		return err
	}
	hc := grpc_health_v1.NewHealthClient(cc)
	maxAttempts := 30
	for {
		resp, err := hc.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
		if err == nil && resp != nil && resp.GetStatus() == grpc_health_v1.HealthCheckResponse_SERVING {
			break
		}
		log.Printf("ATTEMPTING TO CONNECT ")
		maxAttempts--
		if maxAttempts == 0 {
			return errors.New("SUT is not ready for tests")
		}
		<-time.After(time.Duration(1 * time.Second))
	}

	return nil
}

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}
