package bean

type Config struct {
	Section struct {
		Enabled bool
		Debug   bool
		Port    int
	}
	MongoDB struct {
		Ip   string
		Port string
		Database string
	}
}

type User struct {
	Name string `json:"name" bson:"name"`
	Age string `json:"age" bson:"age"`
	Phone string `json:"phone" bason:"phone"`
}

type Res struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data []User `json:"data"`
}