package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迪ArdeeBarony struct {
	feud.BaseBarony
}

var BArdee阿迪 feud.Barony = &阿迪ArdeeBarony{}

func init() {
	f := BArdee阿迪.(*阿迪ArdeeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardee",
		TitleName: "阿迪",
		TitleCode: "b_ardee",
	}
}
