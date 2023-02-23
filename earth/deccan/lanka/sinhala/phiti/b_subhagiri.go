package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 须婆耆厘SubhagiriBarony struct {
	feud.BaseBarony
}

var BSubhagiri须婆耆厘 feud.Barony = &须婆耆厘SubhagiriBarony{}

func init() {
	f := BSubhagiri须婆耆厘.(*须婆耆厘SubhagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subhagiri",
		TitleName: "须婆耆厘",
		TitleCode: "b_subhagiri",
	}
}
