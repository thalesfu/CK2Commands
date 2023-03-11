package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥努特OnutBarony struct {
	feud.BaseBarony
}

var BOnut奥努特 feud.Barony = &奥努特OnutBarony{}

func init() {
    f := BOnut奥努特.(*奥努特OnutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onut",
		TitleName: "奥努特",
		TitleCode: "b_onut",
	}
}
