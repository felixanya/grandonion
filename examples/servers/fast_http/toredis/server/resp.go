package server


type Resp struct {
	Meta 		Meta		`json:"meta"`
	Data 		interface{} `json:"data"`
}

type Meta struct {
	Code 		int		`json:"code"`
	Error 		string	`json:"error"`
}
