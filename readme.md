When a new credential is registered using the registerNewCredential function, the client-side code creates a new
public-private key pair. The private key is kept securely on the user's device and never sent to the server. The public
key, along with some additional metadata, is sent to the server. Here's what you need to store on the server for each
user:

Credential ID: This is a globally unique identifier for this particular credential. It's used to look up the credential
during authentication.  
Public Key: This is used to verify the signature during authentication.  
User Handle: This is the identifier for the user within your system. It could be a user ID, username, email address,
etc.  
Sign Count: This is a counter that the authenticator increments every time it creates an assertion (i.e., every time the
user logs in with this credential). It's used to prevent replay attacks.

WebAuthn is a web standard published by the World Wide Web Consortium (W3C) for passwordless authentication. It allows
servers to register and authenticate users using public key cryptography instead of passwords.

Here are the overall steps involved in registering and authenticating a user using WebAuthn:

**Registration**

1. **Client**: The client requests a new credential from the server. This is typically done when the user is registering
   a new account or adding a new authentication method.

2. **Server**: The server generates a challenge and sends it to the client, along with the user's ID and other data.

3. **Client**: The client calls `navigator.credentials.create()`, passing in the challenge and user data. This causes
   the user's authenticator (e.g., a biometric device or a hardware key) to generate a new public-private key pair.

4. **Client**: The authenticator returns a new credential to the client. This credential includes the public key, the
   user's ID, and a signature over the challenge and some of the data.

5. **Client**: The client sends the new credential back to the server.

6. **Server**: The server verifies the signature on the credential to ensure it's valid and was created by the
   authenticator. If the signature is valid, the server stores the public key and user's ID for later use.

**Authentication**

1. **Client**: The client requests authentication from the server. This is typically done when the user is trying to log
   in.

2. **Server**: The server generates a new challenge and sends it to the client.

3. **Client**: The client calls `navigator.credentials.get()`, passing in the challenge. This causes the user's
   authenticator to sign the challenge with the private key.

4. **Client**: The authenticator returns an assertion to the client. This assertion includes the signature over the
   challenge and some additional data.

5. **Client**: The client sends the assertion back to the server.

6. **Server**: The server verifies the signature on the assertion using the stored public key. If the signature is
   valid, the server knows that the assertion came from the authenticator and that the user is present.

This is a simplified overview of the process. The actual implementation can be more complex and may involve additional
steps, such as checking that the authenticator is trusted and handling errors.

<h3>WebAuthn References</h3>

1. https://github.com/passwordless-id/webauthn?tab=readme-ov-file
2. https://github.com/go-webauthn/webauthn

<h3>TODOs</h3>

1. Refresh token
2. 