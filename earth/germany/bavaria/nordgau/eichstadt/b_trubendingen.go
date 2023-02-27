package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁本丁根TrubendingenBarony struct {
	feud.BaseBarony
}

var BTrubendingen特鲁本丁根 feud.Barony = &特鲁本丁根TrubendingenBarony{}

func init() {
    f := BTrubendingen特鲁本丁根.(*特鲁本丁根TrubendingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trubendingen",
		TitleName: "特鲁本丁根",
		TitleCode: "b_trubendingen",
	}
}
