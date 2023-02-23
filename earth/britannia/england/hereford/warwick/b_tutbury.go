package warwick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔特伯里TutburyBarony struct {
	feud.BaseBarony
}

var BTutbury塔特伯里 feud.Barony = &塔特伯里TutburyBarony{}

func init() {
	f := BTutbury塔特伯里.(*塔特伯里TutburyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tutbury",
		TitleName: "塔特伯里",
		TitleCode: "b_tutbury",
	}
}
