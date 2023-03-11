package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索维拉EssaouiraBarony struct {
	feud.BaseBarony
}

var BEssaouira索维拉 feud.Barony = &索维拉EssaouiraBarony{}

func init() {
    f := BEssaouira索维拉.(*索维拉EssaouiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "essaouira",
		TitleName: "索维拉",
		TitleCode: "b_essaouira",
	}
}
