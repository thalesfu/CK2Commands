package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊利ElyBarony struct {
	feud.BaseBarony
}

var BEly伊利 feud.Barony = &伊利ElyBarony{}

func init() {
    f := BEly伊利.(*伊利ElyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ely",
		TitleName: "伊利",
		TitleCode: "b_ely",
	}
}
