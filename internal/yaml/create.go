package yaml

import (
	"fmt"
	"os"
	"strings"

	"github.com/svang/svangapi/internal/config"
	"gopkg.in/yaml.v2"
)

//CreateSampleConfigFile creates the sample yaml file
func CreateSampleConfigFile(sampleFileName string) error {
	config := &config.SvangType{}
	config.APIVersion = "v1"
	config.Kind = "Configuration"
	config.Metadata.Name = "sVang"
	config.Spec.MAIL.User = "Type email here - vikram@zinox.com"
	config.Spec.MAIL.Password = "Type password here"
	config.Spec.FirebaseConfig.APIKey = "Firebase API Key"
	config.Spec.FirebaseConfig.AppID = "Firebase APP ID"
	config.Spec.FirebaseConfig.AuthDomain = "Firebase Auth Domain"
	config.Spec.FirebaseConfig.DatabaseURL = "Firebase database url"
	config.Spec.FirebaseConfig.MessagingSenderID = "Firevase messasing sender id"
	config.Spec.FirebaseConfig.ProjectID = "Firebase Project ID"
	config.Spec.FirebaseConfig.StorageBucket = "Firebase storage bucket"
	config.Spec.ServiceAccount.Type = "service_account"
	config.Spec.ServiceAccount.TokenURI = "https://oauth2.googleapis.com/token"
	config.Spec.ServiceAccount.ProjectID = "Firebase project id"
	config.Spec.ServiceAccount.PrivateKeyID = "Private key ID"
	config.Spec.ServiceAccount.PrivateKey = "Private Key"
	config.Spec.ServiceAccount.ClientX509CertURL = "x509 cert url"
	config.Spec.ServiceAccount.ClientID = "firebase client id"
	config.Spec.ServiceAccount.ClientEmail = "firebase client email"
	config.Spec.ServiceAccount.AuthURI = "https://accounts.google.com/o/oauth2/auth"
	config.Spec.ServiceAccount.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"

	err := CreateFromConfig(config, sampleFileName)
	if err != nil {
		return fmt.Errorf("error while writing to file %v: %v", sampleFileName, err)
	}
	return nil
}

//CreateFromConfig creates the yaml file from config struct
func CreateFromConfig(config *config.SvangType, filename string) error {

	filename = strings.TrimSpace(filename)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create %v file, err : %v", filename, err)
	}
	content, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error while marshalling data to file %v: %v", filename, err)
	}

	_, err = file.Write(content)

	return err
}
