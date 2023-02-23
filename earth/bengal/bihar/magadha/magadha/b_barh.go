package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗BarhBarony struct {
	feud.BaseBarony
}

var BBarh婆罗 feud.Barony = &婆罗BarhBarony{}

func init() {
	f := BBarh婆罗.(*婆罗BarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barh",
		TitleName: "婆罗",
		TitleCode: "b_barh",
	}
}
