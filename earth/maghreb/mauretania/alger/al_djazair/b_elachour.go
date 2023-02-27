package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿舒尔ElachourBarony struct {
	feud.BaseBarony
}

var BElachour阿舒尔 feud.Barony = &阿舒尔ElachourBarony{}

func init() {
    f := BElachour阿舒尔.(*阿舒尔ElachourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elachour",
		TitleName: "阿舒尔",
		TitleCode: "b_elachour",
	}
}
