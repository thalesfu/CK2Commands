package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕加马PergamonBarony struct {
	feud.BaseBarony
}

var BPergamon帕加马 feud.Barony = &帕加马PergamonBarony{}

func init() {
	f := BPergamon帕加马.(*帕加马PergamonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pergamon",
		TitleName: "帕加马",
		TitleCode: "b_pergamon",
	}
}
