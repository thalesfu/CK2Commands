package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥德索斯OdessosBarony struct {
	feud.BaseBarony
}

var BOdessos奥德索斯 feud.Barony = &奥德索斯OdessosBarony{}

func init() {
    f := BOdessos奥德索斯.(*奥德索斯OdessosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odessos",
		TitleName: "奥德索斯",
		TitleCode: "b_odessos",
	}
}
