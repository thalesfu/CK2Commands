package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查特里斯ChatterisBarony struct {
	feud.BaseBarony
}

var BChatteris查特里斯 feud.Barony = &查特里斯ChatterisBarony{}

func init() {
	f := BChatteris查特里斯.(*查特里斯ChatterisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatteris",
		TitleName: "查特里斯",
		TitleCode: "b_chatteris",
	}
}
