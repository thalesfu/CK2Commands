package chanderi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿诃吉AharjiBarony struct {
	feud.BaseBarony
}

var BAharji阿诃吉 feud.Barony = &阿诃吉AharjiBarony{}

func init() {
    f := BAharji阿诃吉.(*阿诃吉AharjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aharji",
		TitleName: "阿诃吉",
		TitleCode: "b_aharji",
	}
}
