package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉克鲁瓦LacroixBarony struct {
	feud.BaseBarony
}

var BLacroix拉克鲁瓦 feud.Barony = &拉克鲁瓦LacroixBarony{}

func init() {
	f := BLacroix拉克鲁瓦.(*拉克鲁瓦LacroixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lacroix",
		TitleName: "拉克鲁瓦",
		TitleCode: "b_lacroix",
	}
}
