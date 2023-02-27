package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 首罗耆厘SuragiriBarony struct {
	feud.BaseBarony
}

var BSuragiri首罗耆厘 feud.Barony = &首罗耆厘SuragiriBarony{}

func init() {
    f := BSuragiri首罗耆厘.(*首罗耆厘SuragiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suragiri",
		TitleName: "首罗耆厘",
		TitleCode: "b_suragiri",
	}
}
