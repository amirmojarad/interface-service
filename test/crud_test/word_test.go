package crud

import (
	"interface_project/test"
	"testing"
)

func TestWords(t *testing.T) {
	testClient := test.GetTestClientAndContext(t)
	defer testClient.CallCancelAndClose()

}
