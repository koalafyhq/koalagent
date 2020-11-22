package acme

import (
	"context"
	"crypto"
	"io/ioutil"

	"github.com/go-acme/lego/certcrypto"
	"github.com/go-acme/lego/certificate"
	"github.com/go-acme/lego/challenge/http01"
	"github.com/go-acme/lego/lego"
	"github.com/go-acme/lego/registration"
)

// Issuer is
type Issuer struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

// GetEmail is
func (i *Issuer) GetEmail() string {
	return i.Email
}

// GetRegistration is
func (i Issuer) GetRegistration() *registration.Resource {
	return i.Registration
}

// GetPrivateKey is
func (i *Issuer) GetPrivateKey() crypto.PrivateKey {
	return i.key
}

func processGeneratedCertificate(cert string) error {
	// TODO: do something here

	return nil
}

// IssueNewCertificate is
func IssueNewCertificate(ctx context.Context, args ...interface{}) error {
	// FIXME: refactor this
	arg := args[0].(map[string]interface{})

	// TODO: handle when new certificate contains alias like www
	domain := arg["domain"].(string)
	issuerEmail := arg["issuer_email"].(string)
	caDirURL := arg["ca_dir_url"].(string)

	// FIXME: can we do better than this?
	pubKey, err := ioutil.ReadFile("pem/le-public.pem")

	if err != nil {
		Error("acme", err)
	}

	privKey, err := ioutil.ReadFile("pem/le-private.pem")

	if err != nil {
		Error("acme", err)
	}

	accountPrivateKey, _ := Decode(privKey, pubKey)

	if err != nil {
		Error("acme", err)
	}

	issuer := Issuer{
		Email: issuerEmail,
		key:   accountPrivateKey,
	}

	config := lego.NewConfig(&issuer)

	config.CADirURL = caDirURL
	config.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(config)

	if err != nil {
		Error("lego", err)
	}

	err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002"))

	if err != nil {
		Error("lego", err)
	}

	account, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})

	if err != nil {
		Error("lego", err)
	}

	issuer.Registration = account

	request := certificate.ObtainRequest{
		Domains:    []string{domain},
		Bundle:     false,
		MustStaple: true,
	}

	certificates, err := client.Certificate.Obtain(request)

	if err != nil {
		Error("lego", err)
	}

	cert := append(certificates.Certificate, certificates.PrivateKey...)

	Info("acme", "Successfully issuing certificate for "+domain)

	return processGeneratedCertificate(string(cert))
}
