package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车底伽摩ChatigamaBarony struct {
	feud.BaseBarony
}

var BChatigama车底伽摩 feud.Barony = &车底伽摩ChatigamaBarony{}

func init() {
    f := BChatigama车底伽摩.(*车底伽摩ChatigamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatigama",
		TitleName: "车底伽摩",
		TitleCode: "b_chatigama",
	}
}
