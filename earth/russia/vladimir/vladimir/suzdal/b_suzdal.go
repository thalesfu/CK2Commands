package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏兹达尔SuzdalBarony struct {
	feud.BaseBarony
}

var BSuzdal苏兹达尔 feud.Barony = &苏兹达尔SuzdalBarony{}

func init() {
	f := BSuzdal苏兹达尔.(*苏兹达尔SuzdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suzdal",
		TitleName: "苏兹达尔",
		TitleCode: "b_suzdal",
	}
}
