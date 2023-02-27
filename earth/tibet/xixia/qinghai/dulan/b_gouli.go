package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沟里GouliBarony struct {
	feud.BaseBarony
}

var BGouli沟里 feud.Barony = &沟里GouliBarony{}

func init() {
    f := BGouli沟里.(*沟里GouliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gouli",
		TitleName: "沟里",
		TitleCode: "b_gouli",
	}
}
