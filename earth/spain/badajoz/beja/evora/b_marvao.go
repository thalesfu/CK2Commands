package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔旺MarvaoBarony struct {
	feud.BaseBarony
}

var BMarvao马尔旺 feud.Barony = &马尔旺MarvaoBarony{}

func init() {
	f := BMarvao马尔旺.(*马尔旺MarvaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marvao",
		TitleName: "马尔旺",
		TitleCode: "b_marvao",
	}
}
