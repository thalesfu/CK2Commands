package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施维茨SchwyzBarony struct {
	feud.BaseBarony
}

var BSchwyz施维茨 feud.Barony = &施维茨SchwyzBarony{}

func init() {
	f := BSchwyz施维茨.(*施维茨SchwyzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwyz",
		TitleName: "施维茨",
		TitleCode: "b_schwyz",
	}
}
