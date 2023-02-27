package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨扎瓦SazavaBarony struct {
	feud.BaseBarony
}

var BSazava萨扎瓦 feud.Barony = &萨扎瓦SazavaBarony{}

func init() {
    f := BSazava萨扎瓦.(*萨扎瓦SazavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sazava",
		TitleName: "萨扎瓦",
		TitleCode: "b_sazava",
	}
}
