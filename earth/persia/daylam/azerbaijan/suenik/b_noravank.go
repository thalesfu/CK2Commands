package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺拉万克NoravankBarony struct {
	feud.BaseBarony
}

var BNoravank诺拉万克 feud.Barony = &诺拉万克NoravankBarony{}

func init() {
    f := BNoravank诺拉万克.(*诺拉万克NoravankBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noravank",
		TitleName: "诺拉万克",
		TitleCode: "b_noravank",
	}
}
