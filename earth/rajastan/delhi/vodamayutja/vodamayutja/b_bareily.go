package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆雷利BareilyBarony struct {
	feud.BaseBarony
}

var BBareily婆雷利 feud.Barony = &婆雷利BareilyBarony{}

func init() {
    f := BBareily婆雷利.(*婆雷利BareilyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bareily",
		TitleName: "婆雷利",
		TitleCode: "b_bareily",
	}
}
