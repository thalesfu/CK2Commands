package lubusz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯琴KostrzynBarony struct {
	feud.BaseBarony
}

var BKostrzyn科斯琴 feud.Barony = &科斯琴KostrzynBarony{}

func init() {
    f := BKostrzyn科斯琴.(*科斯琴KostrzynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kostrzyn",
		TitleName: "科斯琴",
		TitleCode: "b_kostrzyn",
	}
}
