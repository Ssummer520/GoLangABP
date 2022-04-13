package conf

import (
	"sync"
)
import . "GoLangABP/demo.Core/Model"

var CarInventoryModel CarInventory
var CarLock sync.RWMutex
var Lock sync.RWMutex
var Age = 1
