package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴克法斯特BuckfastBarony struct {
	feud.BaseBarony
}

var BBuckfast巴克法斯特 feud.Barony = &巴克法斯特BuckfastBarony{}

func init() {
    f := BBuckfast巴克法斯特.(*巴克法斯特BuckfastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buckfast",
		TitleName: "巴克法斯特",
		TitleCode: "b_buckfast",
	}
}
