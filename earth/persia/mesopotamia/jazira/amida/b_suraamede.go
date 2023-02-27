package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏拉阿梅德SuraamedeBarony struct {
	feud.BaseBarony
}

var BSuraamede苏拉阿梅德 feud.Barony = &苏拉阿梅德SuraamedeBarony{}

func init() {
    f := BSuraamede苏拉阿梅德.(*苏拉阿梅德SuraamedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suraamede",
		TitleName: "苏拉阿梅德",
		TitleCode: "b_suraamede",
	}
}
