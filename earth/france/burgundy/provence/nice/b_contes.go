package nice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔特ContesBarony struct {
	feud.BaseBarony
}

var BContes孔特 feud.Barony = &孔特ContesBarony{}

func init() {
	f := BContes孔特.(*孔特ContesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "contes",
		TitleName: "孔特",
		TitleCode: "b_contes",
	}
}
