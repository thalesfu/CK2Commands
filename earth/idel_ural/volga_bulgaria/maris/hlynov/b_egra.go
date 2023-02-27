package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶格拉EgraBarony struct {
	feud.BaseBarony
}

var BEgra叶格拉 feud.Barony = &叶格拉EgraBarony{}

func init() {
    f := BEgra叶格拉.(*叶格拉EgraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egra",
		TitleName: "叶格拉",
		TitleCode: "b_egra",
	}
}
