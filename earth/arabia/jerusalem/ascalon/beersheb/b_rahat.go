package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉哈特RahatBarony struct {
	feud.BaseBarony
}

var BRahat拉哈特 feud.Barony = &拉哈特RahatBarony{}

func init() {
    f := BRahat拉哈特.(*拉哈特RahatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rahat",
		TitleName: "拉哈特",
		TitleCode: "b_rahat",
	}
}
