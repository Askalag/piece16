package model

type TimeItem struct {
	Id          int
	Title       string
	Description string
	TimeCost    float32
	TreeLevel   byte
	ParentId    int
}
