package models

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier int    `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string
	UserUUID  string
}
