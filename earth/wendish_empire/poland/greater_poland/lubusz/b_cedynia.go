package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采迪尼亚CedyniaBarony struct {
	feud.BaseBarony
}

var BCedynia采迪尼亚 feud.Barony = &采迪尼亚CedyniaBarony{}

func init() {
    f := BCedynia采迪尼亚.(*采迪尼亚CedyniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cedynia",
		TitleName: "采迪尼亚",
		TitleCode: "b_cedynia",
	}
}
