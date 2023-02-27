package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿科达特AkordatBarony struct {
	feud.BaseBarony
}

var BAkordat阿科达特 feud.Barony = &阿科达特AkordatBarony{}

func init() {
    f := BAkordat阿科达特.(*阿科达特AkordatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akordat",
		TitleName: "阿科达特",
		TitleCode: "b_akordat",
	}
}
