package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基泽乌KizayBarony struct {
	feud.BaseBarony
}

var BKizay基泽乌 feud.Barony = &基泽乌KizayBarony{}

func init() {
	f := BKizay基泽乌.(*基泽乌KizayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizay",
		TitleName: "基泽乌",
		TitleCode: "b_kizay",
	}
}
