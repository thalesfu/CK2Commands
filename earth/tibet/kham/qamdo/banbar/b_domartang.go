package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东马塘DomartangBarony struct {
	feud.BaseBarony
}

var BDomartang东马塘 feud.Barony = &东马塘DomartangBarony{}

func init() {
    f := BDomartang东马塘.(*东马塘DomartangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domartang",
		TitleName: "东马塘",
		TitleCode: "b_domartang",
	}
}
