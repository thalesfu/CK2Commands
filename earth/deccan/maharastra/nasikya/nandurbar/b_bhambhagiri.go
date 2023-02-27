package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梵婆耆厘BhambhagiriBarony struct {
	feud.BaseBarony
}

var BBhambhagiri梵婆耆厘 feud.Barony = &梵婆耆厘BhambhagiriBarony{}

func init() {
    f := BBhambhagiri梵婆耆厘.(*梵婆耆厘BhambhagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhambhagiri",
		TitleName: "梵婆耆厘",
		TitleCode: "b_bhambhagiri",
	}
}
