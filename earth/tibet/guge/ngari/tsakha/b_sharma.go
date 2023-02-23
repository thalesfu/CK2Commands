package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏玛SharmaBarony struct {
	feud.BaseBarony
}

var BSharma夏玛 feud.Barony = &夏玛SharmaBarony{}

func init() {
	f := BSharma夏玛.(*夏玛SharmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharma",
		TitleName: "夏玛",
		TitleCode: "b_sharma",
	}
}
