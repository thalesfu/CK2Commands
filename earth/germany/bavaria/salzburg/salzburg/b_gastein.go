package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加斯泰因GasteinBarony struct {
	feud.BaseBarony
}

var BGastein加斯泰因 feud.Barony = &加斯泰因GasteinBarony{}

func init() {
    f := BGastein加斯泰因.(*加斯泰因GasteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gastein",
		TitleName: "加斯泰因",
		TitleCode: "b_gastein",
	}
}
