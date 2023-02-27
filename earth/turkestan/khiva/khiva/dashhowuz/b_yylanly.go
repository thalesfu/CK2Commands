package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊厄兰雷YylanlyBarony struct {
	feud.BaseBarony
}

var BYylanly伊厄兰雷 feud.Barony = &伊厄兰雷YylanlyBarony{}

func init() {
    f := BYylanly伊厄兰雷.(*伊厄兰雷YylanlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yylanly",
		TitleName: "伊厄兰雷",
		TitleCode: "b_yylanly",
	}
}
