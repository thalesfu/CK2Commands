package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳林NarimBarony struct {
	feud.BaseBarony
}

var BNarim纳林 feud.Barony = &纳林NarimBarony{}

func init() {
    f := BNarim纳林.(*纳林NarimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narim",
		TitleName: "纳林",
		TitleCode: "b_narim",
	}
}
