package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬费拉达PonferradaBarony struct {
	feud.BaseBarony
}

var BPonferrada蓬费拉达 feud.Barony = &蓬费拉达PonferradaBarony{}

func init() {
    f := BPonferrada蓬费拉达.(*蓬费拉达PonferradaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ponferrada",
		TitleName: "蓬费拉达",
		TitleCode: "b_ponferrada",
	}
}
