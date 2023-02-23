package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴库尔BarkurBarony struct {
	feud.BaseBarony
}

var BBarkur巴库尔 feud.Barony = &巴库尔BarkurBarony{}

func init() {
	f := BBarkur巴库尔.(*巴库尔BarkurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barkur",
		TitleName: "巴库尔",
		TitleCode: "b_barkur",
	}
}
