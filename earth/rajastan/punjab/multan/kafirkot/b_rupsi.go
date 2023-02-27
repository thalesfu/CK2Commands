package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢斯RupsiBarony struct {
	feud.BaseBarony
}

var BRupsi卢斯 feud.Barony = &卢斯RupsiBarony{}

func init() {
    f := BRupsi卢斯.(*卢斯RupsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rupsi",
		TitleName: "卢斯",
		TitleCode: "b_rupsi",
	}
}
