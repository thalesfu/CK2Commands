package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪赖哈勒SidirahhalBarony struct {
	feud.BaseBarony
}

var BSidirahhal西迪赖哈勒 feud.Barony = &西迪赖哈勒SidirahhalBarony{}

func init() {
	f := BSidirahhal西迪赖哈勒.(*西迪赖哈勒SidirahhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidirahhal",
		TitleName: "西迪赖哈勒",
		TitleCode: "b_sidirahhal",
	}
}
