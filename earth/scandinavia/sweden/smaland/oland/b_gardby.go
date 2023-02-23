package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈德比GardbyBarony struct {
	feud.BaseBarony
}

var BGardby戈德比 feud.Barony = &戈德比GardbyBarony{}

func init() {
	f := BGardby戈德比.(*戈德比GardbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gardby",
		TitleName: "戈德比",
		TitleCode: "b_gardby",
	}
}
