package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努拉NoraBarony struct {
	feud.BaseBarony
}

var BNora努拉 feud.Barony = &努拉NoraBarony{}

func init() {
	f := BNora努拉.(*努拉NoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nora",
		TitleName: "努拉",
		TitleCode: "b_nora",
	}
}
