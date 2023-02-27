package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比格拉维亚BegarawiyahBarony struct {
	feud.BaseBarony
}

var BBegarawiyah比格拉维亚 feud.Barony = &比格拉维亚BegarawiyahBarony{}

func init() {
    f := BBegarawiyah比格拉维亚.(*比格拉维亚BegarawiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "begarawiyah",
		TitleName: "比格拉维亚",
		TitleCode: "b_begarawiyah",
	}
}
