build:
	protoc -I proto/ proto/farmer/farmer.proto --go_out=plugins=micro:proto

	
	#  GOOS=linux GOARCH=amd64 go build
	#  docker build -t farmer-service .

# run 
#		docker run -p 50051:50051 farmer-service