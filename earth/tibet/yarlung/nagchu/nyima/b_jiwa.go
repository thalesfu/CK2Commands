package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉瓦JiwaBarony struct {
	feud.BaseBarony
}

var BJiwa吉瓦 feud.Barony = &吉瓦JiwaBarony{}

func init() {
    f := BJiwa吉瓦.(*吉瓦JiwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiwa",
		TitleName: "吉瓦",
		TitleCode: "b_jiwa",
	}
}
