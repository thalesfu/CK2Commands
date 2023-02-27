package french_leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯科夫RoscoffBarony struct {
	feud.BaseBarony
}

var BRoscoff罗斯科夫 feud.Barony = &罗斯科夫RoscoffBarony{}

func init() {
    f := BRoscoff罗斯科夫.(*罗斯科夫RoscoffBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roscoff",
		TitleName: "罗斯科夫",
		TitleCode: "b_roscoff",
	}
}
