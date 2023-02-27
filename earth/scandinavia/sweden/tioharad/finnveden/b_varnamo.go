package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦纳穆VarnamoBarony struct {
	feud.BaseBarony
}

var BVarnamo韦纳穆 feud.Barony = &韦纳穆VarnamoBarony{}

func init() {
    f := BVarnamo韦纳穆.(*韦纳穆VarnamoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varnamo",
		TitleName: "韦纳穆",
		TitleCode: "b_varnamo",
	}
}
