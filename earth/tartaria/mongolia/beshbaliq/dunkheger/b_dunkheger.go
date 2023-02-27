package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 冬和格尔DunkhegerBarony struct {
	feud.BaseBarony
}

var BDunkheger冬和格尔 feud.Barony = &冬和格尔DunkhegerBarony{}

func init() {
    f := BDunkheger冬和格尔.(*冬和格尔DunkhegerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunkheger",
		TitleName: "冬和格尔",
		TitleCode: "b_dunkheger",
	}
}
