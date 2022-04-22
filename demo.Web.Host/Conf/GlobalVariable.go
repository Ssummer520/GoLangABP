package Conf

import (
	. "GoLangABP/demo.Core/Model"
	"sync"
)

var CarInventoryModel CarInventory
var CarLock sync.RWMutex
var Lock sync.RWMutex
var Age = 1
