package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德LadeBarony struct {
	feud.BaseBarony
}

var BLade拉德 feud.Barony = &拉德LadeBarony{}

func init() {
    f := BLade拉德.(*拉德LadeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lade",
		TitleName: "拉德",
		TitleCode: "b_lade",
	}
}
