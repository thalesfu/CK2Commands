package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔瓦NarvaBarony struct {
	feud.BaseBarony
}

var BNarva纳尔瓦 feud.Barony = &纳尔瓦NarvaBarony{}

func init() {
    f := BNarva纳尔瓦.(*纳尔瓦NarvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narva",
		TitleName: "纳尔瓦",
		TitleCode: "b_narva",
	}
}
