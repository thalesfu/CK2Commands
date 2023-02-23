package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托迪TodiBarony struct {
	feud.BaseBarony
}

var BTodi托迪 feud.Barony = &托迪TodiBarony{}

func init() {
	f := BTodi托迪.(*托迪TodiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "todi",
		TitleName: "托迪",
		TitleCode: "b_todi",
	}
}
