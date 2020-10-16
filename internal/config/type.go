package config

// SvangType is the main config file needed
type SvangType struct {
	APIVersion string `yaml:"apiVersion" validate:"required" json:"apiVersion"`
	Kind       string `yaml:"kind" validate:"required" json:"kind"`
	Metadata   struct {
		Name string `yaml:"name" json:"name"`
	} `yaml:"metadata" json:"metadata"`

	Spec struct {
		MAIL struct {
			User     string `yaml:"user" validate:"required"`
			Password string `yaml:"password" validate:"required"`
		} `yaml:"mail" json:"mail"`
		ServiceAccount struct {
			Type                    string `yaml:"type" validate:"required" json:"type"`
			ProjectID               string `yaml:"project_id" validate:"required" json:"project_id"`
			PrivateKeyID            string `yaml:"rivate_key_id" validate:"required" json:"rivate_key_id"`
			PrivateKey              string `yaml:"private_key" validate:"required" json:"private_key"`
			ClientEmail             string `yaml:"client_email" validate:"required" json:"client_email"`
			AuthURI                 string `yaml:"client_id" validate:"required" json:"client_id"`
			TokenURI                string `yaml:"auth_uri" validate:"required" json:"auth_uri"`
			AuthProviderX509CertURL string `yaml:"token_uri" validate:"required" json:"token_uri"`
			ClientX509CertURL       string `yaml:"auth_provider_x509_cert_url" validate:"required" json:"auth_provider_x509_cert_url"`
			ClientID                string `yaml:"client_x509_cert_url" validate:"required" json:"client_x509_cert_url"`
		} `yaml:"serviceAccount" json:"serviceAccount"`
		FirebaseConfig struct {
			APIKey            string `yaml:"apiKey" validate:"required" json:"apiKey"`
			AuthDomain        string `yaml:"authDomain" validate:"required" json:"authDomain"`
			DatabaseURL       string `yaml:"databaseURL" validate:"required" json:"databaseURL"`
			ProjectID         string `yaml:"projectId" validate:"required" json:"projectId"`
			StorageBucket     string `yaml:"storageBucket" validate:"required" json:"storageBucket"`
			MessagingSenderID string `yaml:"messagingSenderId" validate:"required" json:"messagingSenderId"`
			AppID             string `yaml:"appId" validate:"required" json:"appId"`
		} `yaml:"firebaseConfig" json:"firebaseConfig"`
	} `yaml:"spec" json:"spec"`
}
