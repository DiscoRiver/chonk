# go-chonky

Tests for hiding/retrieving encrypted data in .PNG image chunks

Current stages I'm working on;

1. Rebuild .png file, and retain byte integrity
2. Append payload to existing byte structure
3. Append payload to imported byte structure (same as above, but doesn't overwrite)
4. Inject payload to pre-determined location within byte structure
5. Inject to random location within byte structure

Payload package - This package is responsible for encrypting/decrypting a payload with an AES cipher.

Extract package - this package is responsible for identifying the payload from an arbitrary position within the given file.