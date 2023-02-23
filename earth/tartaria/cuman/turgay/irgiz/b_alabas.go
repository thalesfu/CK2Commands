package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉巴斯AlabasBarony struct {
	feud.BaseBarony
}

var BAlabas阿拉巴斯 feud.Barony = &阿拉巴斯AlabasBarony{}

func init() {
	f := BAlabas阿拉巴斯.(*阿拉巴斯AlabasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alabas",
		TitleName: "阿拉巴斯",
		TitleCode: "b_alabas",
	}
}
