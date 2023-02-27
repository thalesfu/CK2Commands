package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡德伯莱KondapalliBarony struct {
	feud.BaseBarony
}

var BKondapalli贡德伯莱 feud.Barony = &贡德伯莱KondapalliBarony{}

func init() {
    f := BKondapalli贡德伯莱.(*贡德伯莱KondapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kondapalli",
		TitleName: "贡德伯莱",
		TitleCode: "b_kondapalli",
	}
}
