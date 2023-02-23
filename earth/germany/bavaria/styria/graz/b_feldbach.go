package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔德巴赫FeldbachBarony struct {
	feud.BaseBarony
}

var BFeldbach费尔德巴赫 feud.Barony = &费尔德巴赫FeldbachBarony{}

func init() {
	f := BFeldbach费尔德巴赫.(*费尔德巴赫FeldbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "feldbach",
		TitleName: "费尔德巴赫",
		TitleCode: "b_feldbach",
	}
}
