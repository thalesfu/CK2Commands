package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马朗恩MalangenBarony struct {
	feud.BaseBarony
}

var BMalangen马朗恩 feud.Barony = &马朗恩MalangenBarony{}

func init() {
    f := BMalangen马朗恩.(*马朗恩MalangenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malangen",
		TitleName: "马朗恩",
		TitleCode: "b_malangen",
	}
}
