package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克伊克CuijckBarony struct {
	feud.BaseBarony
}

var BCuijck克伊克 feud.Barony = &克伊克CuijckBarony{}

func init() {
    f := BCuijck克伊克.(*克伊克CuijckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuijck",
		TitleName: "克伊克",
		TitleCode: "b_cuijck",
	}
}
