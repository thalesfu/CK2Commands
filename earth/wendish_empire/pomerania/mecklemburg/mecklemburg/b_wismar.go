package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维斯马WismarBarony struct {
	feud.BaseBarony
}

var BWismar维斯马 feud.Barony = &维斯马WismarBarony{}

func init() {
    f := BWismar维斯马.(*维斯马WismarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wismar",
		TitleName: "维斯马",
		TitleCode: "b_wismar",
	}
}
