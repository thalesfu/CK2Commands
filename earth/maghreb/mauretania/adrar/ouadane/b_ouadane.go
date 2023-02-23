package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦丹OuadaneBarony struct {
	feud.BaseBarony
}

var BOuadane瓦丹 feud.Barony = &瓦丹OuadaneBarony{}

func init() {
	f := BOuadane瓦丹.(*瓦丹OuadaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouadane",
		TitleName: "瓦丹",
		TitleCode: "b_ouadane",
	}
}
