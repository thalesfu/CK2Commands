package ouadane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦丹WadanBarony struct {
	feud.BaseBarony
}

var BWadan瓦丹 feud.Barony = &瓦丹WadanBarony{}

func init() {
    f := BWadan瓦丹.(*瓦丹WadanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wadan",
		TitleName: "瓦丹",
		TitleCode: "b_wadan",
	}
}
