package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎莫希奇ZamoscBarony struct {
	feud.BaseBarony
}

var BZamosc扎莫希奇 feud.Barony = &扎莫希奇ZamoscBarony{}

func init() {
    f := BZamosc扎莫希奇.(*扎莫希奇ZamoscBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zamosc",
		TitleName: "扎莫希奇",
		TitleCode: "b_zamosc",
	}
}
