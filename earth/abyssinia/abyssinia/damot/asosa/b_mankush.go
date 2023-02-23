package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼库什MankushBarony struct {
	feud.BaseBarony
}

var BMankush曼库什 feud.Barony = &曼库什MankushBarony{}

func init() {
	f := BMankush曼库什.(*曼库什MankushBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankush",
		TitleName: "曼库什",
		TitleCode: "b_mankush",
	}
}
