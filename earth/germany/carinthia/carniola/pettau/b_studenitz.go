package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施蒂德尼茨StudenitzBarony struct {
	feud.BaseBarony
}

var BStudenitz施蒂德尼茨 feud.Barony = &施蒂德尼茨StudenitzBarony{}

func init() {
	f := BStudenitz施蒂德尼茨.(*施蒂德尼茨StudenitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "studenitz",
		TitleName: "施蒂德尼茨",
		TitleCode: "b_studenitz",
	}
}
