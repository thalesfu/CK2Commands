package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福伦霍弗VollenhoveBarony struct {
	feud.BaseBarony
}

var BVollenhove福伦霍弗 feud.Barony = &福伦霍弗VollenhoveBarony{}

func init() {
	f := BVollenhove福伦霍弗.(*福伦霍弗VollenhoveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vollenhove",
		TitleName: "福伦霍弗",
		TitleCode: "b_vollenhove",
	}
}
