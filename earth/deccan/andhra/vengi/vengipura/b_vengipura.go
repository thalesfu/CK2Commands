package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓶耆罗VengipuraBarony struct {
	feud.BaseBarony
}

var BVengipura瓶耆罗 feud.Barony = &瓶耆罗VengipuraBarony{}

func init() {
    f := BVengipura瓶耆罗.(*瓶耆罗VengipuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vengipura",
		TitleName: "瓶耆罗",
		TitleCode: "b_vengipura",
	}
}
