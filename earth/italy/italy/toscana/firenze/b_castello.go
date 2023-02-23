package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰洛CastelloBarony struct {
	feud.BaseBarony
}

var BCastello卡斯泰洛 feud.Barony = &卡斯泰洛CastelloBarony{}

func init() {
	f := BCastello卡斯泰洛.(*卡斯泰洛CastelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castello",
		TitleName: "卡斯泰洛",
		TitleCode: "b_castello",
	}
}
