package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦泽莱VezelayBarony struct {
	feud.BaseBarony
}

var BVezelay韦泽莱 feud.Barony = &韦泽莱VezelayBarony{}

func init() {
	f := BVezelay韦泽莱.(*韦泽莱VezelayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vezelay",
		TitleName: "韦泽莱",
		TitleCode: "b_vezelay",
	}
}
