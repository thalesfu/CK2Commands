package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦特WerthBarony struct {
	feud.BaseBarony
}

var BWerth韦特 feud.Barony = &韦特WerthBarony{}

func init() {
    f := BWerth韦特.(*韦特WerthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "werth",
		TitleName: "韦特",
		TitleCode: "b_werth",
	}
}
