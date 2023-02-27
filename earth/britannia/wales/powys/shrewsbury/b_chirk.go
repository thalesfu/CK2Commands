package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彻克ChirkBarony struct {
	feud.BaseBarony
}

var BChirk彻克 feud.Barony = &彻克ChirkBarony{}

func init() {
    f := BChirk彻克.(*彻克ChirkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chirk",
		TitleName: "彻克",
		TitleCode: "b_chirk",
	}
}
