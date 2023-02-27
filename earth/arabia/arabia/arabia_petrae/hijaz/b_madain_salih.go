package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛甸沙勒Madain_salihBarony struct {
	feud.BaseBarony
}

var BMadain_salih玛甸沙勒 feud.Barony = &玛甸沙勒Madain_salihBarony{}

func init() {
    f := BMadain_salih玛甸沙勒.(*玛甸沙勒Madain_salihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madain_salih",
		TitleName: "玛甸沙勒",
		TitleCode: "b_madain_salih",
	}
}
