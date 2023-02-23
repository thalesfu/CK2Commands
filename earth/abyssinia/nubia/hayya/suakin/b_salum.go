package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨鲁姆SalumBarony struct {
	feud.BaseBarony
}

var BSalum萨鲁姆 feud.Barony = &萨鲁姆SalumBarony{}

func init() {
	f := BSalum萨鲁姆.(*萨鲁姆SalumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salum",
		TitleName: "萨鲁姆",
		TitleCode: "b_salum",
	}
}
