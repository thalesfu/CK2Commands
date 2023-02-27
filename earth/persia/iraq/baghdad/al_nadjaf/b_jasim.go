package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾西姆JasimBarony struct {
	feud.BaseBarony
}

var BJasim贾西姆 feud.Barony = &贾西姆JasimBarony{}

func init() {
    f := BJasim贾西姆.(*贾西姆JasimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jasim",
		TitleName: "贾西姆",
		TitleCode: "b_jasim",
	}
}
