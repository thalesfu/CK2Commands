package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔门德拉莱霍AlmendralejoBarony struct {
	feud.BaseBarony
}

var BAlmendralejo阿尔门德拉莱霍 feud.Barony = &阿尔门德拉莱霍AlmendralejoBarony{}

func init() {
	f := BAlmendralejo阿尔门德拉莱霍.(*阿尔门德拉莱霍AlmendralejoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almendralejo",
		TitleName: "阿尔门德拉莱霍",
		TitleCode: "b_almendralejo",
	}
}
