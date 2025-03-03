#!/bin/bash

PI_USER="pi"
PI_IP="192.168.0.130"
PI_DIR="~/piCRT"

echo "🛠️ Building Go binary..."
GOOS=linux GOARCH=arm GOARM=7 go build -o server server.go

echo "🛠️ Building frontend..."
cd piCRT-ui
pnpm run build
cd ..

echo "Syncing files to Pi..."
rsync -avz --delete piCRT-ui/build $PI_USER@$PI_IP:$PI_DIR/
rsync -avz --delete ./server $PI_USER@$PI_IP:$PI_DIR/

echo "Restarting Go server on Pi..."
ssh $PI_USER@$PI_IP "sudo systemctl restart piCRT"

echo "🚀 Deployment complete!"
