package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰尔诺达里CastelnaudaryBarony struct {
	feud.BaseBarony
}

var BCastelnaudary卡斯泰尔诺达里 feud.Barony = &卡斯泰尔诺达里CastelnaudaryBarony{}

func init() {
	f := BCastelnaudary卡斯泰尔诺达里.(*卡斯泰尔诺达里CastelnaudaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelnaudary",
		TitleName: "卡斯泰尔诺达里",
		TitleCode: "b_castelnaudary",
	}
}
