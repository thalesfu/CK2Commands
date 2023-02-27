package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于埃尔戈阿HuelgoatBarony struct {
	feud.BaseBarony
}

var BHuelgoat于埃尔戈阿 feud.Barony = &于埃尔戈阿HuelgoatBarony{}

func init() {
    f := BHuelgoat于埃尔戈阿.(*于埃尔戈阿HuelgoatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huelgoat",
		TitleName: "于埃尔戈阿",
		TitleCode: "b_huelgoat",
	}
}
