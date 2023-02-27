package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富兰陀罗PurandarBarony struct {
	feud.BaseBarony
}

var BPurandar富兰陀罗 feud.Barony = &富兰陀罗PurandarBarony{}

func init() {
    f := BPurandar富兰陀罗.(*富兰陀罗PurandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purandar",
		TitleName: "富兰陀罗",
		TitleCode: "b_purandar",
	}
}
