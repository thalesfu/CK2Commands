package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩他尼耶MathaniaBarony struct {
	feud.BaseBarony
}

var BMathania摩他尼耶 feud.Barony = &摩他尼耶MathaniaBarony{}

func init() {
	f := BMathania摩他尼耶.(*摩他尼耶MathaniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mathania",
		TitleName: "摩他尼耶",
		TitleCode: "b_mathania",
	}
}
