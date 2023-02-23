package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔加雷什GargarescBarony struct {
	feud.BaseBarony
}

var BGargaresc杰尔加雷什 feud.Barony = &杰尔加雷什GargarescBarony{}

func init() {
	f := BGargaresc杰尔加雷什.(*杰尔加雷什GargarescBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gargaresc",
		TitleName: "杰尔加雷什",
		TitleCode: "b_gargaresc",
	}
}
