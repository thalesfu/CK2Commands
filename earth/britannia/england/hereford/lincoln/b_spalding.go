package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯伯丁SpaldingBarony struct {
	feud.BaseBarony
}

var BSpalding斯伯丁 feud.Barony = &斯伯丁SpaldingBarony{}

func init() {
    f := BSpalding斯伯丁.(*斯伯丁SpaldingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spalding",
		TitleName: "斯伯丁",
		TitleCode: "b_spalding",
	}
}
