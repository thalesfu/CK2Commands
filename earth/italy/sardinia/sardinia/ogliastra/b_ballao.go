package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴劳BallaoBarony struct {
	feud.BaseBarony
}

var BBallao巴劳 feud.Barony = &巴劳BallaoBarony{}

func init() {
    f := BBallao巴劳.(*巴劳BallaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ballao",
		TitleName: "巴劳",
		TitleCode: "b_ballao",
	}
}
