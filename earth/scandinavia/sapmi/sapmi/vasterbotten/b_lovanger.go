package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒翁厄尔LovangerBarony struct {
	feud.BaseBarony
}

var BLovanger勒翁厄尔 feud.Barony = &勒翁厄尔LovangerBarony{}

func init() {
    f := BLovanger勒翁厄尔.(*勒翁厄尔LovangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lovanger",
		TitleName: "勒翁厄尔",
		TitleCode: "b_lovanger",
	}
}
