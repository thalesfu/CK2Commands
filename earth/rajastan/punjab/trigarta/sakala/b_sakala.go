package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奢羯罗SakalaBarony struct {
	feud.BaseBarony
}

var BSakala奢羯罗 feud.Barony = &奢羯罗SakalaBarony{}

func init() {
    f := BSakala奢羯罗.(*奢羯罗SakalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakala",
		TitleName: "奢羯罗",
		TitleCode: "b_sakala",
	}
}
