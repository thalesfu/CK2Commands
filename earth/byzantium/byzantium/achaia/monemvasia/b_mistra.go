package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米斯特拉斯MistraBarony struct {
	feud.BaseBarony
}

var BMistra米斯特拉斯 feud.Barony = &米斯特拉斯MistraBarony{}

func init() {
    f := BMistra米斯特拉斯.(*米斯特拉斯MistraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mistra",
		TitleName: "米斯特拉斯",
		TitleCode: "b_mistra",
	}
}
