package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥扎利OzaljBarony struct {
	feud.BaseBarony
}

var BOzalj奥扎利 feud.Barony = &奥扎利OzaljBarony{}

func init() {
	f := BOzalj奥扎利.(*奥扎利OzaljBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozalj",
		TitleName: "奥扎利",
		TitleCode: "b_ozalj",
	}
}
