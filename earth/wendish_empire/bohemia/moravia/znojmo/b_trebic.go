package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特热比奇TrebicBarony struct {
	feud.BaseBarony
}

var BTrebic特热比奇 feud.Barony = &特热比奇TrebicBarony{}

func init() {
    f := BTrebic特热比奇.(*特热比奇TrebicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trebic",
		TitleName: "特热比奇",
		TitleCode: "b_trebic",
	}
}
