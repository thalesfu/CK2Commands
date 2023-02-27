package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗赖施塔特FreistadtBarony struct {
	feud.BaseBarony
}

var BFreistadt弗赖施塔特 feud.Barony = &弗赖施塔特FreistadtBarony{}

func init() {
    f := BFreistadt弗赖施塔特.(*弗赖施塔特FreistadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freistadt",
		TitleName: "弗赖施塔特",
		TitleCode: "b_freistadt",
	}
}
