package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格平根GoppingenBarony struct {
	feud.BaseBarony
}

var BGoppingen格平根 feud.Barony = &格平根GoppingenBarony{}

func init() {
    f := BGoppingen格平根.(*格平根GoppingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goppingen",
		TitleName: "格平根",
		TitleCode: "b_goppingen",
	}
}
