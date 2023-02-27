package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢利ShchellBarony struct {
	feud.BaseBarony
}

var BShchell谢利 feud.Barony = &谢利ShchellBarony{}

func init() {
    f := BShchell谢利.(*谢利ShchellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shchell",
		TitleName: "谢利",
		TitleCode: "b_shchell",
	}
}
