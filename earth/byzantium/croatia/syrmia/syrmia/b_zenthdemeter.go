package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣戴迈泰尔ZenthdemeterBarony struct {
	feud.BaseBarony
}

var BZenthdemeter圣戴迈泰尔 feud.Barony = &圣戴迈泰尔ZenthdemeterBarony{}

func init() {
    f := BZenthdemeter圣戴迈泰尔.(*圣戴迈泰尔ZenthdemeterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zenthdemeter",
		TitleName: "圣戴迈泰尔",
		TitleCode: "b_zenthdemeter",
	}
}
