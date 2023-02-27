package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔卡科尔MarkakolBarony struct {
	feud.BaseBarony
}

var BMarkakol马尔卡科尔 feud.Barony = &马尔卡科尔MarkakolBarony{}

func init() {
    f := BMarkakol马尔卡科尔.(*马尔卡科尔MarkakolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "markakol",
		TitleName: "马尔卡科尔",
		TitleCode: "b_markakol",
	}
}
