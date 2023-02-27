package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞蒂莫SettimoBarony struct {
	feud.BaseBarony
}

var BSettimo塞蒂莫 feud.Barony = &塞蒂莫SettimoBarony{}

func init() {
    f := BSettimo塞蒂莫.(*塞蒂莫SettimoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "settimo",
		TitleName: "塞蒂莫",
		TitleCode: "b_settimo",
	}
}
