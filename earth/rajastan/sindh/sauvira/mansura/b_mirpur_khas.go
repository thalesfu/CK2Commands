package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔布尔哈斯Mirpur_khasBarony struct {
	feud.BaseBarony
}

var BMirpur_khas米尔布尔哈斯 feud.Barony = &米尔布尔哈斯Mirpur_khasBarony{}

func init() {
    f := BMirpur_khas米尔布尔哈斯.(*米尔布尔哈斯Mirpur_khasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirpur_khas",
		TitleName: "米尔布尔哈斯",
		TitleCode: "b_mirpur_khas",
	}
}
