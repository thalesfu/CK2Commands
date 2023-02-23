package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达尔哈米尔BandarkhamirBarony struct {
	feud.BaseBarony
}

var BBandarkhamir班达尔哈米尔 feud.Barony = &班达尔哈米尔BandarkhamirBarony{}

func init() {
	f := BBandarkhamir班达尔哈米尔.(*班达尔哈米尔BandarkhamirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandarkhamir",
		TitleName: "班达尔哈米尔",
		TitleCode: "b_bandarkhamir",
	}
}
