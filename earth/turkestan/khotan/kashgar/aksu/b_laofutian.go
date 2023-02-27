package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 老付田LaofutianBarony struct {
	feud.BaseBarony
}

var BLaofutian老付田 feud.Barony = &老付田LaofutianBarony{}

func init() {
    f := BLaofutian老付田.(*老付田LaofutianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laofutian",
		TitleName: "老付田",
		TitleCode: "b_laofutian",
	}
}
