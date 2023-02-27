package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰尔诺CastelnauBarony struct {
	feud.BaseBarony
}

var BCastelnau卡斯泰尔诺 feud.Barony = &卡斯泰尔诺CastelnauBarony{}

func init() {
    f := BCastelnau卡斯泰尔诺.(*卡斯泰尔诺CastelnauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelnau",
		TitleName: "卡斯泰尔诺",
		TitleCode: "b_castelnau",
	}
}
