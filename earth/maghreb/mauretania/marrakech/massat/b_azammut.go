package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾宰穆尔AzammutBarony struct {
	feud.BaseBarony
}

var BAzammut艾宰穆尔 feud.Barony = &艾宰穆尔AzammutBarony{}

func init() {
    f := BAzammut艾宰穆尔.(*艾宰穆尔AzammutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azammut",
		TitleName: "艾宰穆尔",
		TitleCode: "b_azammut",
	}
}
