package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺加尔特NogalteBarony struct {
	feud.BaseBarony
}

var BNogalte诺加尔特 feud.Barony = &诺加尔特NogalteBarony{}

func init() {
    f := BNogalte诺加尔特.(*诺加尔特NogalteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogalte",
		TitleName: "诺加尔特",
		TitleCode: "b_nogalte",
	}
}
