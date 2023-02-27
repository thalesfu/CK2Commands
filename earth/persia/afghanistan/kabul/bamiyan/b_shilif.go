package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希利夫ShilifBarony struct {
	feud.BaseBarony
}

var BShilif希利夫 feud.Barony = &希利夫ShilifBarony{}

func init() {
    f := BShilif希利夫.(*希利夫ShilifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shilif",
		TitleName: "希利夫",
		TitleCode: "b_shilif",
	}
}
