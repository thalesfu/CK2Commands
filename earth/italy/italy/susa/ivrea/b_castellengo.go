package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰伦戈CastellengoBarony struct {
	feud.BaseBarony
}

var BCastellengo卡斯泰伦戈 feud.Barony = &卡斯泰伦戈CastellengoBarony{}

func init() {
	f := BCastellengo卡斯泰伦戈.(*卡斯泰伦戈CastellengoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellengo",
		TitleName: "卡斯泰伦戈",
		TitleCode: "b_castellengo",
	}
}
