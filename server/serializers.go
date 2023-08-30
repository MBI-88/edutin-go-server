package server

type data struct {
	ID    int `json:"id"`
	Bought bool `json:"bought"`
	Img   string `json:"img"`
	Cate  string `json:"cate"`
	Selected bool `json:"selected"`
}
