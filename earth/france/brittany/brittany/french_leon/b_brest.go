package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷斯特BrestBarony struct {
	feud.BaseBarony
}

var BBrest布雷斯特 feud.Barony = &布雷斯特BrestBarony{}

func init() {
    f := BBrest布雷斯特.(*布雷斯特BrestBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brest",
		TitleName: "布雷斯特",
		TitleCode: "b_brest",
	}
}
