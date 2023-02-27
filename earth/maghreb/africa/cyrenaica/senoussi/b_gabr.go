package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加布罗GabrBarony struct {
	feud.BaseBarony
}

var BGabr加布罗 feud.Barony = &加布罗GabrBarony{}

func init() {
    f := BGabr加布罗.(*加布罗GabrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gabr",
		TitleName: "加布罗",
		TitleCode: "b_gabr",
	}
}
