package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克列梅涅茨KremyanetsBarony struct {
	feud.BaseBarony
}

var BKremyanets克列梅涅茨 feud.Barony = &克列梅涅茨KremyanetsBarony{}

func init() {
    f := BKremyanets克列梅涅茨.(*克列梅涅茨KremyanetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kremyanets",
		TitleName: "克列梅涅茨",
		TitleCode: "b_kremyanets",
	}
}
