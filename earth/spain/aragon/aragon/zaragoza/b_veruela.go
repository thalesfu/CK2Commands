package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦鲁埃拉VeruelaBarony struct {
	feud.BaseBarony
}

var BVeruela韦鲁埃拉 feud.Barony = &韦鲁埃拉VeruelaBarony{}

func init() {
    f := BVeruela韦鲁埃拉.(*韦鲁埃拉VeruelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veruela",
		TitleName: "韦鲁埃拉",
		TitleCode: "b_veruela",
	}
}
