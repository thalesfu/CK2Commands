package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基泰KiteeBarony struct {
	feud.BaseBarony
}

var BKitee基泰 feud.Barony = &基泰KiteeBarony{}

func init() {
    f := BKitee基泰.(*基泰KiteeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kitee",
		TitleName: "基泰",
		TitleCode: "b_kitee",
	}
}
