package middleware

import (
	"fmt"
	"testing"

	http "tacit-api/http"
	"tacit-api/mocks"
	pki "tacit-api/pki"

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

	pkp := &pki.RealPublicKeyProvider{}

	mockLogger.EXPECT().Info("There was an issue validating the JWT: ", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

	//assertions
	if c.SetIsCalled {
		t.Errorf("Set should not be called since the JWT cannot be validated")
	}
}

func TestDoJwtValidationFailsWithoutAKID(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.TCYt5XsITJX1CxPCT8yAV-TVkIEq_PbChOMqsLfRoPsnsgw5WEuts01mq-pQy7UJiN5mgRxD-WUcX16dUEMGlv50aqzpqh4Qktb3rk-BuQy72IFLOqV0G_zS245-kronKb78cPN25DGlcTwLtjPAYuNzVBAh4vGHSrQyHUdBBPM",
	}

	pkp := &pki.RealPublicKeyProvider{}

	mockLogger.EXPECT().Info("There was an issue validating the JWT: ", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

	//assertions
	if c.SetIsCalled {
		t.Errorf("Set should not be called since the JWT cannot be validated")
	}
}

func TestDoJwtValidationHappyPath(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InNvbWVLSUQifQ.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.N_S2f4XzLg_Gz3h7fcuR26I3miiDmDygZwZXXbUMUOOcb-LBBhbrWY-6Op-ArJ-MGQMvJT5QkSyn0sj4Zgnxm8Z2wGG_kuChQjpk9HRx3YMPOBrkLnfh2IR3ovpKCwxXVOBE0kCvQacMcfnif4_OV-mNB-3Pp1sAAdeuIeKUhUo",
	}

	pkp := &pki.RealPublicKeyProvider{}

	mockLogger.EXPECT().Debugf("KID of token %v", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

	fmt.Println("test PASSES")
}
