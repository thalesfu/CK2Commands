package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜赖希德BerrechidBarony struct {
	feud.BaseBarony
}

var BBerrechid拜赖希德 feud.Barony = &拜赖希德BerrechidBarony{}

func init() {
	f := BBerrechid拜赖希德.(*拜赖希德BerrechidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berrechid",
		TitleName: "拜赖希德",
		TitleCode: "b_berrechid",
	}
}
