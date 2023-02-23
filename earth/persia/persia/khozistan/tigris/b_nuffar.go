package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努法尔NuffarBarony struct {
	feud.BaseBarony
}

var BNuffar努法尔 feud.Barony = &努法尔NuffarBarony{}

func init() {
	f := BNuffar努法尔.(*努法尔NuffarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuffar",
		TitleName: "努法尔",
		TitleCode: "b_nuffar",
	}
}
