package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博拉马BoramaBarony struct {
	feud.BaseBarony
}

var BBorama博拉马 feud.Barony = &博拉马BoramaBarony{}

func init() {
    f := BBorama博拉马.(*博拉马BoramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borama",
		TitleName: "博拉马",
		TitleCode: "b_borama",
	}
}
