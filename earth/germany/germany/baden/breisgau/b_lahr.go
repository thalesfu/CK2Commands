package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉尔LahrBarony struct {
	feud.BaseBarony
}

var BLahr拉尔 feud.Barony = &拉尔LahrBarony{}

func init() {
	f := BLahr拉尔.(*拉尔LahrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahr",
		TitleName: "拉尔",
		TitleCode: "b_lahr",
	}
}
