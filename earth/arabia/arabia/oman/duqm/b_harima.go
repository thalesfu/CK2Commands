package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈日玛HarimaBarony struct {
	feud.BaseBarony
}

var BHarima哈日玛 feud.Barony = &哈日玛HarimaBarony{}

func init() {
	f := BHarima哈日玛.(*哈日玛HarimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harima",
		TitleName: "哈日玛",
		TitleCode: "b_harima",
	}
}
