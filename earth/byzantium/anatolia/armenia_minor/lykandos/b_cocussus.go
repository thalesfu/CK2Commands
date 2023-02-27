package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格克孙CocussusBarony struct {
	feud.BaseBarony
}

var BCocussus格克孙 feud.Barony = &格克孙CocussusBarony{}

func init() {
    f := BCocussus格克孙.(*格克孙CocussusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cocussus",
		TitleName: "格克孙",
		TitleCode: "b_cocussus",
	}
}
