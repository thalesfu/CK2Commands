package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗纳德赛RonaldsayBarony struct {
	feud.BaseBarony
}

var BRonaldsay罗纳德赛 feud.Barony = &罗纳德赛RonaldsayBarony{}

func init() {
    f := BRonaldsay罗纳德赛.(*罗纳德赛RonaldsayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ronaldsay",
		TitleName: "罗纳德赛",
		TitleCode: "b_ronaldsay",
	}
}
