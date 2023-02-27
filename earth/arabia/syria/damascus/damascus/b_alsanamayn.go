package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞奈迈因AlsanamaynBarony struct {
	feud.BaseBarony
}

var BAlsanamayn塞奈迈因 feud.Barony = &塞奈迈因AlsanamaynBarony{}

func init() {
    f := BAlsanamayn塞奈迈因.(*塞奈迈因AlsanamaynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsanamayn",
		TitleName: "塞奈迈因",
		TitleCode: "b_alsanamayn",
	}
}
