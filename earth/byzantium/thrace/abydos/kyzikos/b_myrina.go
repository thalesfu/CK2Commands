package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 密里那MyrinaBarony struct {
	feud.BaseBarony
}

var BMyrina密里那 feud.Barony = &密里那MyrinaBarony{}

func init() {
    f := BMyrina密里那.(*密里那MyrinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myrina",
		TitleName: "密里那",
		TitleCode: "b_myrina",
	}
}
