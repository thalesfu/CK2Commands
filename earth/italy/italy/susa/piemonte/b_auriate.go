package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里亚泰AuriateBarony struct {
	feud.BaseBarony
}

var BAuriate奥里亚泰 feud.Barony = &奥里亚泰AuriateBarony{}

func init() {
	f := BAuriate奥里亚泰.(*奥里亚泰AuriateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auriate",
		TitleName: "奥里亚泰",
		TitleCode: "b_auriate",
	}
}
