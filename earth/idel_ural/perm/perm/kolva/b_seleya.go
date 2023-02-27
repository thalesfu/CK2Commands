package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢列亚SeleyaBarony struct {
	feud.BaseBarony
}

var BSeleya谢列亚 feud.Barony = &谢列亚SeleyaBarony{}

func init() {
    f := BSeleya谢列亚.(*谢列亚SeleyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seleya",
		TitleName: "谢列亚",
		TitleCode: "b_seleya",
	}
}
