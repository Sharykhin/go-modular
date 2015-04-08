package tests

import (
		"testing"
		"go-modular/application/controller"
		)


func Test_IncrementMe(t *testing.T) {
	ctrl := new(controller.BaseController)
	if ctrl.IncrementMe(1,4) != 5 {
		t.Error("IncrementMe does't work")
	} else  {
		t.Log("One test passed.")
	}
}


