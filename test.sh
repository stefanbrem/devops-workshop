gcloud kms encrypt \
    --location "asia-southeast1" \
    --keyring "stefan-key-ring" \
    --key "hello-world-k8s" \
    --plaintext-file application.secrets \
    --ciphertext-file application.secrets.encrypted

