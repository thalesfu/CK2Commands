package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基督城Christ_churchBarony struct {
	feud.BaseBarony
}

var BChrist_church基督城 feud.Barony = &基督城Christ_churchBarony{}

func init() {
    f := BChrist_church基督城.(*基督城Christ_churchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "christ_church",
		TitleName: "基督城",
		TitleCode: "b_christ_church",
	}
}
