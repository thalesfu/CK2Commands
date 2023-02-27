package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尤切斯基VyucheyskiyBarony struct {
	feud.BaseBarony
}

var BVyucheyskiy维尤切斯基 feud.Barony = &维尤切斯基VyucheyskiyBarony{}

func init() {
    f := BVyucheyskiy维尤切斯基.(*维尤切斯基VyucheyskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyucheyskiy",
		TitleName: "维尤切斯基",
		TitleCode: "b_vyucheyskiy",
	}
}
