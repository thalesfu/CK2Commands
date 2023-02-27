package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨利巴鲁SalliballuBarony struct {
	feud.BaseBarony
}

var BSalliballu萨利巴鲁 feud.Barony = &萨利巴鲁SalliballuBarony{}

func init() {
    f := BSalliballu萨利巴鲁.(*萨利巴鲁SalliballuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salliballu",
		TitleName: "萨利巴鲁",
		TitleCode: "b_salliballu",
	}
}
