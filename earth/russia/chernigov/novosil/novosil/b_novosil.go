package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺沃西利NovosilBarony struct {
	feud.BaseBarony
}

var BNovosil诺沃西利 feud.Barony = &诺沃西利NovosilBarony{}

func init() {
	f := BNovosil诺沃西利.(*诺沃西利NovosilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novosil",
		TitleName: "诺沃西利",
		TitleCode: "b_novosil",
	}
}
