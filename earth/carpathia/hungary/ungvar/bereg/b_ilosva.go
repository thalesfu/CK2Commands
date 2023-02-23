package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊洛什沃IlosvaBarony struct {
	feud.BaseBarony
}

var BIlosva伊洛什沃 feud.Barony = &伊洛什沃IlosvaBarony{}

func init() {
	f := BIlosva伊洛什沃.(*伊洛什沃IlosvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilosva",
		TitleName: "伊洛什沃",
		TitleCode: "b_ilosva",
	}
}
