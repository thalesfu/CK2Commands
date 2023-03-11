package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈德德拉HaddraaBarony struct {
	feud.BaseBarony
}

var BHaddraa哈德德拉 feud.Barony = &哈德德拉HaddraaBarony{}

func init() {
    f := BHaddraa哈德德拉.(*哈德德拉HaddraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haddraa",
		TitleName: "哈德德拉",
		TitleCode: "b_haddraa",
	}
}
