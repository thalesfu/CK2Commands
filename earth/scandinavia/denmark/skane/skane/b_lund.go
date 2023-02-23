package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆德LundBarony struct {
	feud.BaseBarony
}

var BLund隆德 feud.Barony = &隆德LundBarony{}

func init() {
	f := BLund隆德.(*隆德LundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lund",
		TitleName: "隆德",
		TitleCode: "b_lund",
	}
}
