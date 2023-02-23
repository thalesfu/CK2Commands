package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔菲MelfiBarony struct {
	feud.BaseBarony
}

var BMelfi梅尔菲 feud.Barony = &梅尔菲MelfiBarony{}

func init() {
	f := BMelfi梅尔菲.(*梅尔菲MelfiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melfi",
		TitleName: "梅尔菲",
		TitleCode: "b_melfi",
	}
}
