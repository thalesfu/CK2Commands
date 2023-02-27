package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普拉克西诺ApraksinoBarony struct {
	feud.BaseBarony
}

var BApraksino阿普拉克西诺 feud.Barony = &阿普拉克西诺ApraksinoBarony{}

func init() {
    f := BApraksino阿普拉克西诺.(*阿普拉克西诺ApraksinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apraksino",
		TitleName: "阿普拉克西诺",
		TitleCode: "b_apraksino",
	}
}
