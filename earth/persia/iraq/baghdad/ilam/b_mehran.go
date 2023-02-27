package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅赫兰MehranBarony struct {
	feud.BaseBarony
}

var BMehran梅赫兰 feud.Barony = &梅赫兰MehranBarony{}

func init() {
    f := BMehran梅赫兰.(*梅赫兰MehranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehran",
		TitleName: "梅赫兰",
		TitleCode: "b_mehran",
	}
}
