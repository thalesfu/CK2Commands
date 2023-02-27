package racakonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 部婆那耆厘BhuvanagiriBarony struct {
	feud.BaseBarony
}

var BBhuvanagiri部婆那耆厘 feud.Barony = &部婆那耆厘BhuvanagiriBarony{}

func init() {
    f := BBhuvanagiri部婆那耆厘.(*部婆那耆厘BhuvanagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhuvanagiri",
		TitleName: "部婆那耆厘",
		TitleCode: "b_bhuvanagiri",
	}
}
