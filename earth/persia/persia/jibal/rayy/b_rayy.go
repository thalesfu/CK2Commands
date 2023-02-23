package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷伊RayyBarony struct {
	feud.BaseBarony
}

var BRayy雷伊 feud.Barony = &雷伊RayyBarony{}

func init() {
	f := BRayy雷伊.(*雷伊RayyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayy",
		TitleName: "雷伊",
		TitleCode: "b_rayy",
	}
}
