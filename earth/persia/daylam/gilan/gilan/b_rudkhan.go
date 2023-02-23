package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁德罕RudkhanBarony struct {
	feud.BaseBarony
}

var BRudkhan鲁德罕 feud.Barony = &鲁德罕RudkhanBarony{}

func init() {
	f := BRudkhan鲁德罕.(*鲁德罕RudkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rudkhan",
		TitleName: "鲁德罕",
		TitleCode: "b_rudkhan",
	}
}
