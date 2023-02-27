package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门巴MambaBarony struct {
	feud.BaseBarony
}

var BMamba门巴 feud.Barony = &门巴MambaBarony{}

func init() {
    f := BMamba门巴.(*门巴MambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamba",
		TitleName: "门巴",
		TitleCode: "b_mamba",
	}
}
