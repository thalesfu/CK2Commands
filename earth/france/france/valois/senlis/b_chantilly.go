package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚蒂伊ChantillyBarony struct {
	feud.BaseBarony
}

var BChantilly尚蒂伊 feud.Barony = &尚蒂伊ChantillyBarony{}

func init() {
    f := BChantilly尚蒂伊.(*尚蒂伊ChantillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chantilly",
		TitleName: "尚蒂伊",
		TitleCode: "b_chantilly",
	}
}
