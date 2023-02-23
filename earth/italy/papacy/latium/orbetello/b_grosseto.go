package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格罗塞托GrossetoBarony struct {
	feud.BaseBarony
}

var BGrosseto格罗塞托 feud.Barony = &格罗塞托GrossetoBarony{}

func init() {
	f := BGrosseto格罗塞托.(*格罗塞托GrossetoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grosseto",
		TitleName: "格罗塞托",
		TitleCode: "b_grosseto",
	}
}
