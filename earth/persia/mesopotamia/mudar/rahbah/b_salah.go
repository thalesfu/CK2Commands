package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉赫SalahBarony struct {
	feud.BaseBarony
}

var BSalah萨拉赫 feud.Barony = &萨拉赫SalahBarony{}

func init() {
    f := BSalah萨拉赫.(*萨拉赫SalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salah",
		TitleName: "萨拉赫",
		TitleCode: "b_salah",
	}
}
