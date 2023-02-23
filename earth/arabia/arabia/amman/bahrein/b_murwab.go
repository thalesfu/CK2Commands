package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆尔瓦比MurwabBarony struct {
	feud.BaseBarony
}

var BMurwab穆尔瓦比 feud.Barony = &穆尔瓦比MurwabBarony{}

func init() {
	f := BMurwab穆尔瓦比.(*穆尔瓦比MurwabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murwab",
		TitleName: "穆尔瓦比",
		TitleCode: "b_murwab",
	}
}
