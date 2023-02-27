package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马阿迪MaadiBarony struct {
	feud.BaseBarony
}

var BMaadi马阿迪 feud.Barony = &马阿迪MaadiBarony{}

func init() {
    f := BMaadi马阿迪.(*马阿迪MaadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maadi",
		TitleName: "马阿迪",
		TitleCode: "b_maadi",
	}
}
