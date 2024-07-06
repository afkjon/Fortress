#/bin/bash

echo "---------------------------------------------------------"
echo "Attempting to automatically setup tailscale connection..."
echo "---------------------------------------------------------"

sudo tailscale up --accept-routes && tailscale configure kubeconfig tailscale-operator