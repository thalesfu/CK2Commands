package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗KaraBarony struct {
	feud.BaseBarony
}

var BKara迦罗 feud.Barony = &迦罗KaraBarony{}

func init() {
	f := BKara迦罗.(*迦罗KaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara",
		TitleName: "迦罗",
		TitleCode: "b_kara",
	}
}
