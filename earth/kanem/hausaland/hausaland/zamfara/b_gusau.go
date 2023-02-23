package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古绍GusauBarony struct {
	feud.BaseBarony
}

var BGusau古绍 feud.Barony = &古绍GusauBarony{}

func init() {
	f := BGusau古绍.(*古绍GusauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gusau",
		TitleName: "古绍",
		TitleCode: "b_gusau",
	}
}
