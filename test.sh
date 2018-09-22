mkdir -p cover || echo ''
go test -coverpkg=./... -coverprofile=./cover/coverage.data -timeout=5s ./...
go tool cover -html=./cover/coverage.data -o ./cover/coverage.html
go tool cover -func=./cover/coverage.data -o./ cover/coverage.txt

open ./cover/coverage.html