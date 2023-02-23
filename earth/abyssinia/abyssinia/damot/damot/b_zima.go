package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济马ZimaBarony struct {
	feud.BaseBarony
}

var BZima济马 feud.Barony = &济马ZimaBarony{}

func init() {
	f := BZima济马.(*济马ZimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zima",
		TitleName: "济马",
		TitleCode: "b_zima",
	}
}
