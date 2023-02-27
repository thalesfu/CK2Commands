package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱里亚AleriaBarony struct {
	feud.BaseBarony
}

var BAleria阿莱里亚 feud.Barony = &阿莱里亚AleriaBarony{}

func init() {
    f := BAleria阿莱里亚.(*阿莱里亚AleriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aleria",
		TitleName: "阿莱里亚",
		TitleCode: "b_aleria",
	}
}
