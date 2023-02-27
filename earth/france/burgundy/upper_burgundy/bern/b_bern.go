package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯尔尼BernBarony struct {
	feud.BaseBarony
}

var BBern伯尔尼 feud.Barony = &伯尔尼BernBarony{}

func init() {
    f := BBern伯尔尼.(*伯尔尼BernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bern",
		TitleName: "伯尔尼",
		TitleCode: "b_bern",
	}
}
