package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮海湖PyhaselkaBarony struct {
	feud.BaseBarony
}

var BPyhaselka皮海湖 feud.Barony = &皮海湖PyhaselkaBarony{}

func init() {
    f := BPyhaselka皮海湖.(*皮海湖PyhaselkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyhaselka",
		TitleName: "皮海湖",
		TitleCode: "b_pyhaselka",
	}
}
