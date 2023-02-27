package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿博赫尔AboharBarony struct {
	feud.BaseBarony
}

var BAbohar阿博赫尔 feud.Barony = &阿博赫尔AboharBarony{}

func init() {
    f := BAbohar阿博赫尔.(*阿博赫尔AboharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abohar",
		TitleName: "阿博赫尔",
		TitleCode: "b_abohar",
	}
}
