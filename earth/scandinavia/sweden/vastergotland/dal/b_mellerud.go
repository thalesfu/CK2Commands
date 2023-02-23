package dal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅勒鲁MellerudBarony struct {
	feud.BaseBarony
}

var BMellerud梅勒鲁 feud.Barony = &梅勒鲁MellerudBarony{}

func init() {
	f := BMellerud梅勒鲁.(*梅勒鲁MellerudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mellerud",
		TitleName: "梅勒鲁",
		TitleCode: "b_mellerud",
	}
}
