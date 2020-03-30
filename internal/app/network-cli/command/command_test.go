package command

import (
	"fmt"
	"strings"
	"testing"
)

const host string = "hedgehogregistry.co.uk"

func TestNS(t *testing.T) {
	nss, err := NS(host)
	if err != nil {
		t.Error(err)
	}

	var expected strings.Builder
	expected.WriteString(fmt.Sprintln("ns1.hedgehogregistry.co.uk."))
	expected.WriteString(fmt.Sprintln("ns2.hedgehogregistry.co.uk."))
	if nss != expected.String() {
		t.Errorf("NS(%s) failed. expected %s, got %v", host, expected.String(), nss)
	}
}
