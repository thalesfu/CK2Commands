package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺沃肖尔基NovoselkiBarony struct {
	feud.BaseBarony
}

var BNovoselki诺沃肖尔基 feud.Barony = &诺沃肖尔基NovoselkiBarony{}

func init() {
	f := BNovoselki诺沃肖尔基.(*诺沃肖尔基NovoselkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novoselki",
		TitleName: "诺沃肖尔基",
		TitleCode: "b_novoselki",
	}
}
