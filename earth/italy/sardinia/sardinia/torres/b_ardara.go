package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔达拉ArdaraBarony struct {
	feud.BaseBarony
}

var BArdara阿尔达拉 feud.Barony = &阿尔达拉ArdaraBarony{}

func init() {
	f := BArdara阿尔达拉.(*阿尔达拉ArdaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardara",
		TitleName: "阿尔达拉",
		TitleCode: "b_ardara",
	}
}
