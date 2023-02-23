package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙阿基ShakiBarony struct {
	feud.BaseBarony
}

var BShaki沙阿基 feud.Barony = &沙阿基ShakiBarony{}

func init() {
	f := BShaki沙阿基.(*沙阿基ShakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaki",
		TitleName: "沙阿基",
		TitleCode: "b_shaki",
	}
}
