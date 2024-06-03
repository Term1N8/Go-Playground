package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(certificates)
	certificates.PersistentFlags().Bool("generate", false, "Create new certificates for mTLS server/client.")
	certificates.PersistentFlags().Bool("show", false, "View certificates.")

}

var certificates = &cobra.Command{
	Use:   "certificates",
	Short: "oh geez rick, guess we gotta do something with certificates huh?",
	Long:  `oh geez rick, guess we gotta do something with certificates huh?`,
	Run: func(cmd *cobra.Command, args []string) {
		generateBool, _ := (cmd.Flags().GetBool("generate"))
		if generateBool {
			fmt.Println("Generating new certificates for mTLS")
			key, _ := rsa.GenerateKey(rand.Reader, 4096)
			keyBytes := x509.MarshalPKCS1PrivateKey(key)
			// PEM encoding of private key
			keyPEM := pem.EncodeToMemory(
				&pem.Block{
					Type:  "RSA PRIVATE KEY",
					Bytes: keyBytes,
				},
			)

			os.WriteFile("resources/key.pem", []byte(keyPEM), 0644)

			notBefore := time.Now()
			notAfter := notBefore.Add(365 * 24 * 10 * time.Hour)

			//Create certificate templet
			template := x509.Certificate{
				SerialNumber:          big.NewInt(0),
				Subject:               pkix.Name{CommonName: "localhost"},
				SignatureAlgorithm:    x509.SHA256WithRSA,
				NotBefore:             notBefore,
				NotAfter:              notAfter,
				BasicConstraintsValid: true,
				KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment,
				ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
			}
			//Create certificate using templet
			derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)

			//pem encoding of certificate
			certPem := string(pem.EncodeToMemory(
				&pem.Block{
					Type:  "CERTIFICATE",
					Bytes: derBytes,
				},
			))

			os.WriteFile("resources/cert.pem", []byte(certPem), 0644)
		}

	},
}
