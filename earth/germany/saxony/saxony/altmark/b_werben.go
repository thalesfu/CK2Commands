package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦本WerbenBarony struct {
	feud.BaseBarony
}

var BWerben韦本 feud.Barony = &韦本WerbenBarony{}

func init() {
	f := BWerben韦本.(*韦本WerbenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "werben",
		TitleName: "韦本",
		TitleCode: "b_werben",
	}
}
