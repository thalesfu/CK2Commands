package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尼亚热瓦茨KnjazevacBarony struct {
	feud.BaseBarony
}

var BKnjazevac克尼亚热瓦茨 feud.Barony = &克尼亚热瓦茨KnjazevacBarony{}

func init() {
	f := BKnjazevac克尼亚热瓦茨.(*克尼亚热瓦茨KnjazevacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knjazevac",
		TitleName: "克尼亚热瓦茨",
		TitleCode: "b_knjazevac",
	}
}
