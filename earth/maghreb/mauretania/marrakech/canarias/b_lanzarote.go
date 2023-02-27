package canarias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰萨罗特LanzaroteBarony struct {
	feud.BaseBarony
}

var BLanzarote兰萨罗特 feud.Barony = &兰萨罗特LanzaroteBarony{}

func init() {
    f := BLanzarote兰萨罗特.(*兰萨罗特LanzaroteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanzarote",
		TitleName: "兰萨罗特",
		TitleCode: "b_lanzarote",
	}
}
