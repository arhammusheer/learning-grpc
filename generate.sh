#!/bin/bash

USER_SERVICE_DIR="user-service/proto"
MESSAGE_SERVICE_DIR="message-service/proto"
BACKUP_DIRS=()

# Make backup of directories if they exist
if [ -d "$USER_SERVICE_DIR" ]; then
    mv $USER_SERVICE_DIR "${USER_SERVICE_DIR}.bak"
    BACKUP_DIRS+=("${USER_SERVICE_DIR}.bak")
fi

if [ -d "$MESSAGE_SERVICE_DIR" ]; then
    mv $MESSAGE_SERVICE_DIR "${MESSAGE_SERVICE_DIR}.bak"
    BACKUP_DIRS+=("${MESSAGE_SERVICE_DIR}.bak")
fi

# Cleanup function to restore from backup in case of failure
cleanup() {
    echo "An error occurred. Restoring from backups..."
    for BACKUP in "${BACKUP_DIRS[@]}"; do
        ORIGINAL_DIR="${BACKUP%.bak}"
        if [ -d "$BACKUP" ]; then
            rm -rf "$ORIGINAL_DIR"
            mv "$BACKUP" "$ORIGINAL_DIR"
            echo "Restored $ORIGINAL_DIR from backup."
        fi
    done
    exit 1
}

# Set trap to call cleanup on any error
trap cleanup ERR

# Create directories
mkdir -p "$USER_SERVICE_DIR"
mkdir -p "$MESSAGE_SERVICE_DIR"

# Generate code for User Service
protoc --go_out="$USER_SERVICE_DIR" --go-grpc_out="$USER_SERVICE_DIR" proto/user.proto

# Generate code for Message Service
protoc --go_out="$MESSAGE_SERVICE_DIR" --go-grpc_out="$MESSAGE_SERVICE_DIR" proto/message.proto 


# Move the generated files to the correct directory
mv message-service/proto/proto/message/* message-service/proto/
rm -rf message-service/proto/proto

mv user-service/proto/proto/user/* user-service/proto/
rm -rf user-service/proto/proto


# Remove backups if everything succeeds
for BACKUP in "${BACKUP_DIRS[@]}"; do
    rm -rf "$BACKUP"
    echo "Removed backup $BACKUP as generation was successful."
done


echo "Protobuf generation completed successfully."
