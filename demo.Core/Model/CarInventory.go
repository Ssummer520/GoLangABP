package Model

import "container/list"

type CarInventory struct {
	List       list.List //εΊε­ιε
	SellCount  int
	TotalCount int
	Rate       float64
	X          float64
}
