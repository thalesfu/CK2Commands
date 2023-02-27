package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉布加AlabugaBarony struct {
	feud.BaseBarony
}

var BAlabuga阿拉布加 feud.Barony = &阿拉布加AlabugaBarony{}

func init() {
    f := BAlabuga阿拉布加.(*阿拉布加AlabugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alabuga",
		TitleName: "阿拉布加",
		TitleCode: "b_alabuga",
	}
}
