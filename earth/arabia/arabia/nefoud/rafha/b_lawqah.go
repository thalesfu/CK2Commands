package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳盖LawqahBarony struct {
	feud.BaseBarony
}

var BLawqah劳盖 feud.Barony = &劳盖LawqahBarony{}

func init() {
	f := BLawqah劳盖.(*劳盖LawqahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lawqah",
		TitleName: "劳盖",
		TitleCode: "b_lawqah",
	}
}
