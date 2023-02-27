package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔巴德MalbardeBarony struct {
	feud.BaseBarony
}

var BMalbarde马尔巴德 feud.Barony = &马尔巴德MalbardeBarony{}

func init() {
    f := BMalbarde马尔巴德.(*马尔巴德MalbardeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malbarde",
		TitleName: "马尔巴德",
		TitleCode: "b_malbarde",
	}
}
