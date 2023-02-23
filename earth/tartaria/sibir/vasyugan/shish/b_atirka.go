package shish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿季尔卡AtirkaBarony struct {
	feud.BaseBarony
}

var BAtirka阿季尔卡 feud.Barony = &阿季尔卡AtirkaBarony{}

func init() {
	f := BAtirka阿季尔卡.(*阿季尔卡AtirkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atirka",
		TitleName: "阿季尔卡",
		TitleCode: "b_atirka",
	}
}
