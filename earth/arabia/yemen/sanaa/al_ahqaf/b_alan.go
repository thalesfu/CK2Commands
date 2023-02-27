package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兰AlanBarony struct {
	feud.BaseBarony
}

var BAlan阿兰 feud.Barony = &阿兰AlanBarony{}

func init() {
    f := BAlan阿兰.(*阿兰AlanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alan",
		TitleName: "阿兰",
		TitleCode: "b_alan",
	}
}
