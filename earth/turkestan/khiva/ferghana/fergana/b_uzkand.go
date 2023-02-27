package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讹迹邗UzkandBarony struct {
	feud.BaseBarony
}

var BUzkand讹迹邗 feud.Barony = &讹迹邗UzkandBarony{}

func init() {
    f := BUzkand讹迹邗.(*讹迹邗UzkandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uzkand",
		TitleName: "讹迹邗",
		TitleCode: "b_uzkand",
	}
}
