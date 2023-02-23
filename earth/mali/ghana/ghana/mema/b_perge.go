package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尔热PergeBarony struct {
	feud.BaseBarony
}

var BPerge佩尔热 feud.Barony = &佩尔热PergeBarony{}

func init() {
	f := BPerge佩尔热.(*佩尔热PergeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perge",
		TitleName: "佩尔热",
		TitleCode: "b_perge",
	}
}
