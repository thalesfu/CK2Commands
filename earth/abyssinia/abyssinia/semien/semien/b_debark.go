package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德巴尔格DebarkBarony struct {
	feud.BaseBarony
}

var BDebark德巴尔格 feud.Barony = &德巴尔格DebarkBarony{}

func init() {
	f := BDebark德巴尔格.(*德巴尔格DebarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debark",
		TitleName: "德巴尔格",
		TitleCode: "b_debark",
	}
}
