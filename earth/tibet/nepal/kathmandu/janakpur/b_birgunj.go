package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗罗健阇BirgunjBarony struct {
	feud.BaseBarony
}

var BBirgunj毗罗健阇 feud.Barony = &毗罗健阇BirgunjBarony{}

func init() {
	f := BBirgunj毗罗健阇.(*毗罗健阇BirgunjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birgunj",
		TitleName: "毗罗健阇",
		TitleCode: "b_birgunj",
	}
}
