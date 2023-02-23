package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考斯CowesBarony struct {
	feud.BaseBarony
}

var BCowes考斯 feud.Barony = &考斯CowesBarony{}

func init() {
	f := BCowes考斯.(*考斯CowesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cowes",
		TitleName: "考斯",
		TitleCode: "b_cowes",
	}
}
