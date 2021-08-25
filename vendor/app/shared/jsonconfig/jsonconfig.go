package jsonconfig

import (
	"log"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file
func Load(configFile string, p Parser) {

	jsonBytes := []byte(`{"Database":{"Type":"Bolt","Bolt":{"Path":"gowebapp.db"}},"Email":{"Username":"","Password":"","Hostname":"","Port":25,"From":""},"Recaptcha":{"Enabled":false,"Secret":"","SiteKey":""},"Server":{"Hostname":"","UseHTTP":true,"UseHTTPS":false,"HTTPPort":8080,"HTTPSPort":443,"CertFile":"tls/server.crt","KeyFile":"tls/server.key"},"Session":{"SecretKey":"@r4B?EThaSEh_drudR7P_hub=s#s2Pah","Name":"gosess","Options":{"Path":"/","Domain":"","MaxAge":28800,"Secure":false,"HttpOnly":true}},"Template":{"Root":"base","Children":["partial/menu","partial/footer"]},"View":{"BaseURI":"/","Extension":"tmpl","Folder":"template","Name":"blank","Caching":true}}`)
	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
