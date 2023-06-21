package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/iyarkov/foundation/support"
	"math/big"
	"os"
	"time"
)

var (
	serviceName, dnsName = "admin_client", "localhost"
)

func main() {
	//initCA(
	newKey(serviceName, dnsName)
	signCsr(serviceName)
}

func initCA() {
	fmt.Println("Initializing Certificate Authority")
	if checkIfExists("generated/ca_key.pem") {
		panic("CA Key already exist")
	}
	if checkIfExists("generated/ca_cert.pem") {
		panic("CA Certificate already exist")
	}
	caPrivKey := generatePrivKey("generated/ca_key.pem")
	caCert := generateCertificate("Certificate Authority", "generated/ca_cert.pem", &caPrivKey.PublicKey, nil, caPrivKey, "")
	fmt.Printf("Done, %s\n", caCert.Subject)
}

func newKey(serviceName string, dnsName string) {
	fmt.Println("Initializing Certificate Authority")
	privateKey := generatePrivKey(fmt.Sprintf("generated/%s_key.pem", serviceName))
	generateSignRequest(serviceName, dnsName, privateKey)
}

func checkIfExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}

func generatePrivKey(keyFileName string) *ecdsa.PrivateKey {
	var privKey *ecdsa.PrivateKey
	fmt.Printf("Generating new %s key\n", keyFileName)
	var err error
	privKey, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}

	keyBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		panic(err)
	}
	keyFile, err := os.Create(keyFileName)
	if err != nil {
		panic(err)
	}
	defer support.CloseWithWarning(context.Background(), keyFile, fmt.Sprintf("can not close private key file %s", keyFileName))

	if err := pem.Encode(keyFile, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBytes,
	}); err != nil {
		panic(err)
	}
	return privKey
}

func loadPrivKey(keyFileName string) *ecdsa.PrivateKey {
	fmt.Printf("reading Private Key from %s\n", keyFileName)
	encodedBytes, err := os.ReadFile(keyFileName)
	if err != nil {
		panic(err)
	}
	pemBlock, _ := pem.Decode(encodedBytes)
	privKey, err := x509.ParseECPrivateKey(pemBlock.Bytes)
	if err != nil {
		panic(err)
	}
	return privKey
}

func generateCertificate(appName string, certFileName string, pubKey *ecdsa.PublicKey, caCert *x509.Certificate, caPrivKey *ecdsa.PrivateKey, dnsName string) *x509.Certificate {
	var cert *x509.Certificate
	fmt.Printf("generating new  certificate, writing into %s\n", certFileName)
	isCA := caCert == nil
	cert = &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			CommonName:         appName,
			OrganizationalUnit: []string{"Foundation"},
			Organization:       []string{"Foo Bar, LLC"},
			Country:            []string{"US"},
			Province:           []string{"XY"},
			Locality:           []string{"Fakeville"},
			StreetAddress:      []string{"12345 Main St"},
			PostalCode:         []string{"11111"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  isCA,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		DNSNames:              []string{dnsName},
		BasicConstraintsValid: true,
	}
	if caCert == nil {
		caCert = cert
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, pubKey, caPrivKey)
	if err != nil {
		panic(err)
	}
	certFile, err := os.Create(certFileName)
	if err != nil {
		panic(err)
	}
	defer support.CloseWithWarning(context.Background(), certFile, fmt.Sprintf("can not close certificate file %s", certFileName))

	err = pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		panic(err)
	}
	return cert
}

func loadCertificate(certFileName string) *x509.Certificate {
	fmt.Printf("reading  certificate from %s\n", certFileName)
	encodedBytes, err := os.ReadFile(certFileName)
	if err != nil {
		panic(err)
	}
	pemBlock, _ := pem.Decode(encodedBytes)
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		panic(err)
	}
	return cert
}

func generateSignRequest(serviceName string, dnsName string, key *ecdsa.PrivateKey) {
	subj := pkix.Name{
		CommonName:         serviceName,
		OrganizationalUnit: []string{"Foundation"},
		Organization:       []string{"Foo Bar, LLC"},
		Country:            []string{"US"},
		Province:           []string{"XY"},
		Locality:           []string{"Fakeville"},
		StreetAddress:      []string{"12345 Main St"},
		PostalCode:         []string{"11111"},
	}

	template := x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.ECDSAWithSHA512,
		DNSNames:           []string{dnsName},
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
	if err != nil {
		panic(err)
	}
	certFileName := fmt.Sprintf("generated/%s_csr.pem", serviceName)
	csrFile, err := os.Create(certFileName)
	if err != nil {
		panic(err)
	}
	defer support.CloseWithWarning(context.Background(), csrFile, fmt.Sprintf("can not close certificate file %s", certFileName))
	err = pem.Encode(csrFile, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})
	if err != nil {
		panic(err)
	}

}

func signCsr(serviceName string) {
	caCert := loadCertificate("generated/ca_cert.pem")
	caPrivKey := loadPrivKey("generated/ca_key.pem")
	csr := loadCsr(fmt.Sprintf("generated/%s_csr.pem", serviceName))
	certFileName := fmt.Sprintf("generated/%s_cert.pem", serviceName)
	servicePubKey, ok := csr.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("unsupported public key %v", csr.PublicKey))
	}
	generateCertificate(serviceName, certFileName, servicePubKey, caCert, caPrivKey, csr.DNSNames[0])

}

func loadCsr(fileName string) *x509.CertificateRequest {
	fmt.Printf("reading certificate sign request from %s\n", fileName)
	encodedBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	pemBlock, _ := pem.Decode(encodedBytes)
	csr, err := x509.ParseCertificateRequest(pemBlock.Bytes)
	if err != nil {
		panic(err)
	}
	return csr
}
