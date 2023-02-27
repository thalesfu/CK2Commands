package tavda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尔YarBarony struct {
	feud.BaseBarony
}

var BYar亚尔 feud.Barony = &亚尔YarBarony{}

func init() {
    f := BYar亚尔.(*亚尔YarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yar",
		TitleName: "亚尔",
		TitleCode: "b_yar",
	}
}
