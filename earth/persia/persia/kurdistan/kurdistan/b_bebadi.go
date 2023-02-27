package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝巴迪BebadiBarony struct {
	feud.BaseBarony
}

var BBebadi贝巴迪 feud.Barony = &贝巴迪BebadiBarony{}

func init() {
    f := BBebadi贝巴迪.(*贝巴迪BebadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bebadi",
		TitleName: "贝巴迪",
		TitleCode: "b_bebadi",
	}
}
