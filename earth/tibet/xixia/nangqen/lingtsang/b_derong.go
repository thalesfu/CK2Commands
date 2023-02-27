package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德荣马DerongBarony struct {
	feud.BaseBarony
}

var BDerong德荣马 feud.Barony = &德荣马DerongBarony{}

func init() {
    f := BDerong德荣马.(*德荣马DerongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derong",
		TitleName: "德荣马",
		TitleCode: "b_derong",
	}
}
