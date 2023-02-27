package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加鲁梅勒GaroumeleBarony struct {
	feud.BaseBarony
}

var BGaroumele加鲁梅勒 feud.Barony = &加鲁梅勒GaroumeleBarony{}

func init() {
    f := BGaroumele加鲁梅勒.(*加鲁梅勒GaroumeleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garoumele",
		TitleName: "加鲁梅勒",
		TitleCode: "b_garoumele",
	}
}
