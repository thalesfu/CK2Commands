package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿难多耆厘AnanthagiriBarony struct {
	feud.BaseBarony
}

var BAnanthagiri阿难多耆厘 feud.Barony = &阿难多耆厘AnanthagiriBarony{}

func init() {
    f := BAnanthagiri阿难多耆厘.(*阿难多耆厘AnanthagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ananthagiri",
		TitleName: "阿难多耆厘",
		TitleCode: "b_ananthagiri",
	}
}
