package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾宰穆尔AzemmourBarony struct {
	feud.BaseBarony
}

var BAzemmour艾宰穆尔 feud.Barony = &艾宰穆尔AzemmourBarony{}

func init() {
	f := BAzemmour艾宰穆尔.(*艾宰穆尔AzemmourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azemmour",
		TitleName: "艾宰穆尔",
		TitleCode: "b_azemmour",
	}
}
