package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣保罗SanpawlBarony struct {
	feud.BaseBarony
}

var BSanpawl圣保罗 feud.Barony = &圣保罗SanpawlBarony{}

func init() {
	f := BSanpawl圣保罗.(*圣保罗SanpawlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanpawl",
		TitleName: "圣保罗",
		TitleCode: "b_sanpawl",
	}
}
