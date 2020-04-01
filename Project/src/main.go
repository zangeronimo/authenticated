package main

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
)

//SigningMethod Implement SigningMethod to add new methods for signing or verifying tokens.
type SigningMethod interface {
	Verify(signingString, signature string, key interface{}) error // Returns nil if signature is valid
	Sign(signingString string, key interface{}) (string, error)    // Returns encoded signature or error
	Alg() string                                                   // returns the alg identifier for this method (example: 'HS256')
}

//Token A JWT Token. Different fields will be used depending on whether you're creating or parsing/verifying a token.
type Token struct {
	Raw       string                 // The raw token.  Populated when you Parse a token
	Method    SigningMethod          // The signing method used or to be used
	Header    map[string]interface{} // The first segment of the token
	Claims    Claims                 // The second segment of the token
	Signature string                 // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}

//Claims For a type to be a Claims object, it must just have a Valid method that determines if the token is invalid for any supported reason
type Claims interface {
	Valid() error
}

//Keyfunc Parse methods use this callback function to supply the key for verification. The function receives the parsed, but unverified Token. This allows you to use properties in the Header of the token (such as `kid`) to identify which key to use.
type Keyfunc func(*Token) (interface{}, error)

//MapClaims Claims type that uses the map[string]interface{} for JSON decoding This is the default claims type if you don't supply one
type MapClaims map[string]interface{}

//Valid Validates time based claims "exp, iat, nbf". There is no accounting for clock skew. As well, if any of the above claims are not in the token, it will still be considered a valid claim.
func (m MapClaims) Valid() error {
	return nil
}

func main() {
	http.HandleFunc("/auth", helloServer)
	http.ListenAndServe(":4001", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//DecodeSegment Decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	return nil, nil
}

//EncodeSegment Encode JWT specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
	return ""
}

//ParseECPrivateKeyFromPEM Parse PEM encoded Elliptic Curve Private Key Structure
func ParseECPrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	return &ecdsa.PrivateKey{}, nil
}

//ParseECPublicKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 public key
func ParseECPublicKeyFromPEM(key []byte) (*ecdsa.PublicKey, error) {
	return &ecdsa.PublicKey{}, nil
}

//RegisterSigningMethod Register the "alg" name and a factory function for signing method. This is typically done during init() in the method's implementation
func RegisterSigningMethod(alg string, f func() SigningMethod) {

}
