package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶利瓦勒GallivareBarony struct {
	feud.BaseBarony
}

var BGallivare耶利瓦勒 feud.Barony = &耶利瓦勒GallivareBarony{}

func init() {
    f := BGallivare耶利瓦勒.(*耶利瓦勒GallivareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gallivare",
		TitleName: "耶利瓦勒",
		TitleCode: "b_gallivare",
	}
}
