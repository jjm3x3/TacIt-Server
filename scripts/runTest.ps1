go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

New-Item mocks -ErrorAction Ignore -Type Directory

..\..\bin\mockgen.exe -destination mocks/mock_db.go -package mocks tacit-api/db TacitDB
..\..\bin\mockgen.exe -destination mocks/mock_logger.go -package mocks github.com/sirupsen/logrus FieldLogger

go test 
