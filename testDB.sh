#!/bin/bash

start_test_db() {
  docker run --name hex-arch-testing-db -p 5433:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=template1 -d postgres
}

stop_test_db() {
  docker stop hex-arch-testing-db
  docker rm hex-arch-testing-db
}

while getopts t: flag; do
  case "${flag}" in
    t) type=${OPTARG} ;;
  esac
done

case $type in
  start)
    start_test_db
    echo "Test database started."
    ;;
  unit)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Unit ./...
    stop_test_db
    echo "Test database stopped."
    ;;
  integration)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out -run Integration ./...
    stop_test_db
    echo "Test database stopped."
    ;;
  *)
    start_test_db
    CGO_ENABLED=0 go test -v -p 1 -count=1 -covermode=count -coverprofile=coverage/c.out ./...
    stop_test_db
    echo "Test database stopped."
    ;;
esac
