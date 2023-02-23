package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯佩格AspergBarony struct {
	feud.BaseBarony
}

var BAsperg阿斯佩格 feud.Barony = &阿斯佩格AspergBarony{}

func init() {
	f := BAsperg阿斯佩格.(*阿斯佩格AspergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asperg",
		TitleName: "阿斯佩格",
		TitleCode: "b_asperg",
	}
}
