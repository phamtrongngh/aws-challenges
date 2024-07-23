# KMS Basic Challenge

In this challenge, we will get familiar with the AWS Key Management Service (KMS) and how to use it to encrypt and decrypt data. We will also learn how to use the AWS CLI to encrypt and decrypt data using KMS.

## Prerequisites
- AWS CLI installed and configured with the necessary credentials to interact with AWS services.

## Overview
We will create a symmetric KMS key and use it to encrypt and decrypt the provided plain text file `plain.txt`. 

## Scenario


## Steps
### 1. Create a symmetric KMS key
```bash
export KEY_ID=$(aws kms create-key \
--description "KMS - Basic challenge key" \
--key-spec SYMMETRIC_DEFAULT \
--key-usage ENCRYPT_DECRYPT \
--query KeyMetadata.KeyId \
--output text)
```


### 2. (Optional) Create alias for the KMS key for easy reference
Because the KMS key ID is an UUID string so it is not user-friendly and hard to remember. So we can create an alias for it.
```bash
export KEY_ALIAS=alias/kms-basic-challenge
```

```bash
aws kms create-alias \
--alias-name ${KEY_ALIAS} \
--target-key-id $KEY_ID
```
Now we can use the alias instead of the key ID value. 

### 3. Encrypt the plain text file `plain.txt` using the KMS key
Now content of the `plain.txt` is `THIS IS A PLAIN TEXT FILE, WE NEED TO ENCRYPT IT`. \
We will encrypt it using the KMS key created in step 1:
```bash
aws kms encrypt \
--key-id $KEY_ID \
--plaintext fileb://plain.txt \
--output text \
--query CiphertextBlob | base64 --decode > encrypted.txt
```
The encrypted ouput text is returned in base64 format so we use `base64 --decode` to decode it and then save it to `encrypted.txt`.

### 4. Decrypt the encrypted text file `encrypted.txt` using the KMS key
```bash
aws kms decrypt \
--ciphertext-blob fileb://encrypted.txt \
--output text \
--query Plaintext | base64 --decode > decrypted.txt
```
The decrypted output text is returned in base64 format so we use `base64 --decode` to decode it and then save it to `decrypted.txt`. \
Now we can verify it. If the decrypted text matches the original text, the challenge is completed successfully.

### 5. Clean up
```bash
aws kms schedule-key-deletion \
--key-id $KEY_ID \
--pending-window-in-days 7
```