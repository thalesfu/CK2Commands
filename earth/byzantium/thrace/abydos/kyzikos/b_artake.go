package kyzikos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔代基ArtakeBarony struct {
	feud.BaseBarony
}

var BArtake阿尔代基 feud.Barony = &阿尔代基ArtakeBarony{}

func init() {
	f := BArtake阿尔代基.(*阿尔代基ArtakeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artake",
		TitleName: "阿尔代基",
		TitleCode: "b_artake",
	}
}
