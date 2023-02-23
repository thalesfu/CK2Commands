package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瑟伦LatheronBarony struct {
	feud.BaseBarony
}

var BLatheron拉瑟伦 feud.Barony = &拉瑟伦LatheronBarony{}

func init() {
	f := BLatheron拉瑟伦.(*拉瑟伦LatheronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latheron",
		TitleName: "拉瑟伦",
		TitleCode: "b_latheron",
	}
}
