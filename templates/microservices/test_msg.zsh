#!/bin/zsh

# Set the server address and port
SERVER_ADDRESS="localhost:50051"

# Set the service and method names
SERVICE_NAME="microservices.YourService"
METHOD_NAME="HandleMethod"

# Create a JSON representation of the request
REQUEST_JSON='{"name": "World"}'

# Function to make the gRPC call
make_grpc_call() {
    grpcurl -plaintext -d "${REQUEST_JSON}" \
        ${SERVER_ADDRESS} \
        ${SERVICE_NAME}/${METHOD_NAME}
}

# Main loop
while true; do
    echo "Pinging gRPC server..."
    if output=$(make_grpc_call 2>&1); then
        echo "Server responded: ${output}"
    else
        echo "Error: ${output}"
    fi
    echo "------------------------"
    sleep 5
done