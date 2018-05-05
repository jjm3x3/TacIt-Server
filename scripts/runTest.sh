#!/bin/sh
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

mkdir -p mocks 

mockgen -destination mocks/mock_db.go -package mocks tacit-api/db TacitDB
mockgen -destination mocks/mock_logger.go -package mocks github.com/sirupsen/logrus FieldLogger

go test 
