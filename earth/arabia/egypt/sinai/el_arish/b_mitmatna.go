package el_arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米特马纳MitmatnaBarony struct {
	feud.BaseBarony
}

var BMitmatna米特马纳 feud.Barony = &米特马纳MitmatnaBarony{}

func init() {
    f := BMitmatna米特马纳.(*米特马纳MitmatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mitmatna",
		TitleName: "米特马纳",
		TitleCode: "b_mitmatna",
	}
}
