FROM alpine:latest

run makdir /app
workdir /app

add farmer-service /app/farmer-service

cmd ["./farmer-service"]