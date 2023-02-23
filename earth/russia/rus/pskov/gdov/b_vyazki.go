package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚兹基VyazkiBarony struct {
	feud.BaseBarony
}

var BVyazki维亚兹基 feud.Barony = &维亚兹基VyazkiBarony{}

func init() {
	f := BVyazki维亚兹基.(*维亚兹基VyazkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyazki",
		TitleName: "维亚兹基",
		TitleCode: "b_vyazki",
	}
}
