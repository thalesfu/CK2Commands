package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃东尼AydounBarony struct {
	feud.BaseBarony
}

var BAydoun埃东尼 feud.Barony = &埃东尼AydounBarony{}

func init() {
	f := BAydoun埃东尼.(*埃东尼AydounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aydoun",
		TitleName: "埃东尼",
		TitleCode: "b_aydoun",
	}
}
