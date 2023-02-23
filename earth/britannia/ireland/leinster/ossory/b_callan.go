package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡伦CallanBarony struct {
	feud.BaseBarony
}

var BCallan卡伦 feud.Barony = &卡伦CallanBarony{}

func init() {
	f := BCallan卡伦.(*卡伦CallanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "callan",
		TitleName: "卡伦",
		TitleCode: "b_callan",
	}
}
