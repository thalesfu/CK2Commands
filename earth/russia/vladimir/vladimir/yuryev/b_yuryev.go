package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤里耶夫YuryevBarony struct {
	feud.BaseBarony
}

var BYuryev尤里耶夫 feud.Barony = &尤里耶夫YuryevBarony{}

func init() {
    f := BYuryev尤里耶夫.(*尤里耶夫YuryevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuryev",
		TitleName: "尤里耶夫",
		TitleCode: "b_yuryev",
	}
}
