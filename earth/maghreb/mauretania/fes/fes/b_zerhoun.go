package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎霍ZerhounBarony struct {
	feud.BaseBarony
}

var BZerhoun扎霍 feud.Barony = &扎霍ZerhounBarony{}

func init() {
	f := BZerhoun扎霍.(*扎霍ZerhounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zerhoun",
		TitleName: "扎霍",
		TitleCode: "b_zerhoun",
	}
}
