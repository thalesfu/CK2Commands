package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建诃耆厘KanhagiriBarony struct {
	feud.BaseBarony
}

var BKanhagiri建诃耆厘 feud.Barony = &建诃耆厘KanhagiriBarony{}

func init() {
	f := BKanhagiri建诃耆厘.(*建诃耆厘KanhagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanhagiri",
		TitleName: "建诃耆厘",
		TitleCode: "b_kanhagiri",
	}
}
