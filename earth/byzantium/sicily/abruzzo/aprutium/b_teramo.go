package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰拉莫TeramoBarony struct {
	feud.BaseBarony
}

var BTeramo泰拉莫 feud.Barony = &泰拉莫TeramoBarony{}

func init() {
	f := BTeramo泰拉莫.(*泰拉莫TeramoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teramo",
		TitleName: "泰拉莫",
		TitleCode: "b_teramo",
	}
}
