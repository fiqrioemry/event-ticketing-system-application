#!/bin/bash

# TICKETING APP DEPLOYMENT SCRIPT

echo "🚀 Deploying Ticketing App Server..."
docker-compose -p ticketing_app down -v

echo "🚀 Build container ...."
docker-compose -p ticketing_app up -d --build

echo "✅ Deployment complete!"
