package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫鲁别舒夫HrubeshivBarony struct {
	feud.BaseBarony
}

var BHrubeshiv赫鲁别舒夫 feud.Barony = &赫鲁别舒夫HrubeshivBarony{}

func init() {
    f := BHrubeshiv赫鲁别舒夫.(*赫鲁别舒夫HrubeshivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hrubeshiv",
		TitleName: "赫鲁别舒夫",
		TitleCode: "b_hrubeshiv",
	}
}
