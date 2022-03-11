package v1

//Auth Property
type AuthProperty struct {
	ClusterUuid string `json:"cluster_uuid,omitempty"` //cluster uuid
	ClientUuid  string `json:"client_uuid,omitempty"`  //client uuid
	Assertion   string `json:"assertion,omitempty"`    //<bearer-token>
	// GrantType   string `json:"grant_type,omitempty" default:"urn:ietf:params:oauth:grant-type:jwt-bearer"` //grant_type
}

//Auth
type Auth struct {
	AuthProperty `json:",inline"` //inline property
}

//HTTP REQUEST BODY: Auth
type HttpReqAuth struct {
	Auth `json:",inline"`
}