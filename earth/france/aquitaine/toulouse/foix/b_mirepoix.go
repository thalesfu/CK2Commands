package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔普瓦MirepoixBarony struct {
	feud.BaseBarony
}

var BMirepoix米尔普瓦 feud.Barony = &米尔普瓦MirepoixBarony{}

func init() {
    f := BMirepoix米尔普瓦.(*米尔普瓦MirepoixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirepoix",
		TitleName: "米尔普瓦",
		TitleCode: "b_mirepoix",
	}
}
