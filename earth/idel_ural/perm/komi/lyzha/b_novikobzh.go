package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺维科布日NovikobzhBarony struct {
	feud.BaseBarony
}

var BNovikobzh诺维科布日 feud.Barony = &诺维科布日NovikobzhBarony{}

func init() {
    f := BNovikobzh诺维科布日.(*诺维科布日NovikobzhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novikobzh",
		TitleName: "诺维科布日",
		TitleCode: "b_novikobzh",
	}
}
