package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣多纳茨St_donatsBarony struct {
	feud.BaseBarony
}

var BSt_donats圣多纳茨 feud.Barony = &圣多纳茨St_donatsBarony{}

func init() {
    f := BSt_donats圣多纳茨.(*圣多纳茨St_donatsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "st_donats",
		TitleName: "圣多纳茨",
		TitleCode: "b_st_donats",
	}
}
