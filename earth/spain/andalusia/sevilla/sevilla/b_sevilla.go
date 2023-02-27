package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞维利亚SevillaBarony struct {
	feud.BaseBarony
}

var BSevilla塞维利亚 feud.Barony = &塞维利亚SevillaBarony{}

func init() {
    f := BSevilla塞维利亚.(*塞维利亚SevillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sevilla",
		TitleName: "塞维利亚",
		TitleCode: "b_sevilla",
	}
}
