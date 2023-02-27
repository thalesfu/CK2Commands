package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥穆特AmotBarony struct {
	feud.BaseBarony
}

var BAmot奥穆特 feud.Barony = &奥穆特AmotBarony{}

func init() {
    f := BAmot奥穆特.(*奥穆特AmotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amot",
		TitleName: "奥穆特",
		TitleCode: "b_amot",
	}
}
