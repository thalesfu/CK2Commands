package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯特基UrtsekiBarony struct {
	feud.BaseBarony
}

var BUrtseki乌斯特基 feud.Barony = &乌斯特基UrtsekiBarony{}

func init() {
    f := BUrtseki乌斯特基.(*乌斯特基UrtsekiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urtseki",
		TitleName: "乌斯特基",
		TitleCode: "b_urtseki",
	}
}
