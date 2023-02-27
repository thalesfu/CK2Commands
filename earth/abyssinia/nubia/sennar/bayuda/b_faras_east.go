package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东法拉斯Faras_eastBarony struct {
	feud.BaseBarony
}

var BFaras_east东法拉斯 feud.Barony = &东法拉斯Faras_eastBarony{}

func init() {
    f := BFaras_east东法拉斯.(*东法拉斯Faras_eastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faras_east",
		TitleName: "东法拉斯",
		TitleCode: "b_faras_east",
	}
}
