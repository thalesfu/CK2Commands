package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔博勒查KombolchaBarony struct {
	feud.BaseBarony
}

var BKombolcha孔博勒查 feud.Barony = &孔博勒查KombolchaBarony{}

func init() {
    f := BKombolcha孔博勒查.(*孔博勒查KombolchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kombolcha",
		TitleName: "孔博勒查",
		TitleCode: "b_kombolcha",
	}
}
