# FaktoidAPI
Specification, reference implementation and tools for sharing brief location-based facts with the world 

## Running
```
env GOOS=linux GOARCH=amd64 go build src/rahvafakt/svc/RahvaSvc.go
zip ../rsvc.zip RahvaSvc EHAK2015v1.txt RV0241_utf.csv Dockerfile Dockerrun.aws.json 
```
The resulting bundle can be deployed at any docker-capable platform
