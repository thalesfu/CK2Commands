package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文蒂米利亚VentimigliaBarony struct {
	feud.BaseBarony
}

var BVentimiglia文蒂米利亚 feud.Barony = &文蒂米利亚VentimigliaBarony{}

func init() {
    f := BVentimiglia文蒂米利亚.(*文蒂米利亚VentimigliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ventimiglia",
		TitleName: "文蒂米利亚",
		TitleCode: "b_ventimiglia",
	}
}
