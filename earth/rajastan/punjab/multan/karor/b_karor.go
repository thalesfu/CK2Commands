package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦卢罗KarorBarony struct {
	feud.BaseBarony
}

var BKaror迦卢罗 feud.Barony = &迦卢罗KarorBarony{}

func init() {
	f := BKaror迦卢罗.(*迦卢罗KarorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karor",
		TitleName: "迦卢罗",
		TitleCode: "b_karor",
	}
}
