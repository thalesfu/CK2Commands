package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰克伦堡TecklenburgBarony struct {
	feud.BaseBarony
}

var BTecklenburg泰克伦堡 feud.Barony = &泰克伦堡TecklenburgBarony{}

func init() {
    f := BTecklenburg泰克伦堡.(*泰克伦堡TecklenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tecklenburg",
		TitleName: "泰克伦堡",
		TitleCode: "b_tecklenburg",
	}
}
