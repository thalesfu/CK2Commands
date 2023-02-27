package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆耆厘DevagiriBarony struct {
	feud.BaseBarony
}

var BDevagiri提婆耆厘 feud.Barony = &提婆耆厘DevagiriBarony{}

func init() {
    f := BDevagiri提婆耆厘.(*提婆耆厘DevagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devagiri",
		TitleName: "提婆耆厘",
		TitleCode: "b_devagiri",
	}
}
