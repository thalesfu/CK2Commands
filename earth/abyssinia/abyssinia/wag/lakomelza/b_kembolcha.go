package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔博勒查KembolchaBarony struct {
	feud.BaseBarony
}

var BKembolcha孔博勒查 feud.Barony = &孔博勒查KembolchaBarony{}

func init() {
    f := BKembolcha孔博勒查.(*孔博勒查KembolchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kembolcha",
		TitleName: "孔博勒查",
		TitleCode: "b_kembolcha",
	}
}
