package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜格阿BaqaBarony struct {
	feud.BaseBarony
}

var BBaqa拜格阿 feud.Barony = &拜格阿BaqaBarony{}

func init() {
    f := BBaqa拜格阿.(*拜格阿BaqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baqa",
		TitleName: "拜格阿",
		TitleCode: "b_baqa",
	}
}
