package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那地耶陀NadiadBarony struct {
	feud.BaseBarony
}

var BNadiad那地耶陀 feud.Barony = &那地耶陀NadiadBarony{}

func init() {
    f := BNadiad那地耶陀.(*那地耶陀NadiadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nadiad",
		TitleName: "那地耶陀",
		TitleCode: "b_nadiad",
	}
}
