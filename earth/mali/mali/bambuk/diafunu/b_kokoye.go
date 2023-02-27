package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科科耶KokoyeBarony struct {
	feud.BaseBarony
}

var BKokoye科科耶 feud.Barony = &科科耶KokoyeBarony{}

func init() {
    f := BKokoye科科耶.(*科科耶KokoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokoye",
		TitleName: "科科耶",
		TitleCode: "b_kokoye",
	}
}
