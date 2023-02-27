package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梯罗达TirourdaBarony struct {
	feud.BaseBarony
}

var BTirourda梯罗达 feud.Barony = &梯罗达TirourdaBarony{}

func init() {
    f := BTirourda梯罗达.(*梯罗达TirourdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirourda",
		TitleName: "梯罗达",
		TitleCode: "b_tirourda",
	}
}
