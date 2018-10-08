package middleware

import (
	"fmt"
	"testing"

	http "tacit-api/http"
	"tacit-api/mocks"

	"github.com/golang/mock/gomock"
)

func TestDoJwtValidationFailsWithBadJwt(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: "NotEnough.segments",
	}

	mockLogger.EXPECT().Info("There was an issue validating the JWT: ", gomock.Any())

	// execution
	doJwtValidation(mockLogger, c)

	//taken care of in mock expectations
}

func TestDoJwtValidationHappyPath(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InNvbWVLSUQifQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.N_S2f4XzLg_Gz3h7fcuR26I3miiDmDygZwZXXbUMUOOcb-LBBhbrWY-6Op-ArJ-MGQMvJT5QkSyn0sj4Zgnxm8Z2wGG_kuChQjpk9HRx3YMPOBrkLnfh2IR3ovpKCwxXVOBE0kCvQacMcfnif4_OV-mNB-3Pp1sAAdeuIeKUhUo",
	}

	mockLogger.EXPECT().Debugf("KID of token %v", gomock.Any())

	// execution
	doJwtValidation(mockLogger, c)

	fmt.Println("test PASSES")
}
