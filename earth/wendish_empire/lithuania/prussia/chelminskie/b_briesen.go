package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里森BriesenBarony struct {
	feud.BaseBarony
}

var BBriesen布里森 feud.Barony = &布里森BriesenBarony{}

func init() {
    f := BBriesen布里森.(*布里森BriesenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "briesen",
		TitleName: "布里森",
		TitleCode: "b_briesen",
	}
}
