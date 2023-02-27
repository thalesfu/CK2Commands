package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比亚瓦采尔凯夫BialocerkevBarony struct {
	feud.BaseBarony
}

var BBialocerkev比亚瓦采尔凯夫 feud.Barony = &比亚瓦采尔凯夫BialocerkevBarony{}

func init() {
    f := BBialocerkev比亚瓦采尔凯夫.(*比亚瓦采尔凯夫BialocerkevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bialocerkev",
		TitleName: "比亚瓦采尔凯夫",
		TitleCode: "b_bialocerkev",
	}
}
