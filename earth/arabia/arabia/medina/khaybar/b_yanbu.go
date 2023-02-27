package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延布YanbuBarony struct {
	feud.BaseBarony
}

var BYanbu延布 feud.Barony = &延布YanbuBarony{}

func init() {
    f := BYanbu延布.(*延布YanbuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yanbu",
		TitleName: "延布",
		TitleCode: "b_yanbu",
	}
}
