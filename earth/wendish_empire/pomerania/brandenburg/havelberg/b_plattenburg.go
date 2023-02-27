package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉滕堡PlattenburgBarony struct {
	feud.BaseBarony
}

var BPlattenburg普拉滕堡 feud.Barony = &普拉滕堡PlattenburgBarony{}

func init() {
    f := BPlattenburg普拉滕堡.(*普拉滕堡PlattenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plattenburg",
		TitleName: "普拉滕堡",
		TitleCode: "b_plattenburg",
	}
}
