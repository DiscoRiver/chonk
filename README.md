# got-chonk

Tests for hiding encrypted data in .PNG images

Current stages I'm working on;

1. Appending Chunk to the image
2. Inserting chunk to a pre-determined location.

Aside from the technology to insert chunks into .PNG files, there are other stages to completing this project;

1. To encrypt data using AES cipher
2. Insert that data into the image
3. Transporting file non-destructively (compression affects image bytes)
3. Be able to locate and decrypt data on remote end.