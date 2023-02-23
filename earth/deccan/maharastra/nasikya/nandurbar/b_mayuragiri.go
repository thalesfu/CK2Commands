package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩由罗耆厘MayuragiriBarony struct {
	feud.BaseBarony
}

var BMayuragiri摩由罗耆厘 feud.Barony = &摩由罗耆厘MayuragiriBarony{}

func init() {
	f := BMayuragiri摩由罗耆厘.(*摩由罗耆厘MayuragiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayuragiri",
		TitleName: "摩由罗耆厘",
		TitleCode: "b_mayuragiri",
	}
}
