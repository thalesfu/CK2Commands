package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔盖TurgayBarony struct {
	feud.BaseBarony
}

var BTurgay图尔盖 feud.Barony = &图尔盖TurgayBarony{}

func init() {
    f := BTurgay图尔盖.(*图尔盖TurgayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turgay",
		TitleName: "图尔盖",
		TitleCode: "b_turgay",
	}
}
