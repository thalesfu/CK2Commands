package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏卜瓦ShabwaBarony struct {
	feud.BaseBarony
}

var BShabwa夏卜瓦 feud.Barony = &夏卜瓦ShabwaBarony{}

func init() {
	f := BShabwa夏卜瓦.(*夏卜瓦ShabwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shabwa",
		TitleName: "夏卜瓦",
		TitleCode: "b_shabwa",
	}
}
