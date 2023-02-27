package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡莫AlcarnoBarony struct {
	feud.BaseBarony
}

var BAlcarno阿尔卡莫 feud.Barony = &阿尔卡莫AlcarnoBarony{}

func init() {
    f := BAlcarno阿尔卡莫.(*阿尔卡莫AlcarnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcarno",
		TitleName: "阿尔卡莫",
		TitleCode: "b_alcarno",
	}
}
