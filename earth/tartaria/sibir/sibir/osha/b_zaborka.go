package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎博尔卡ZaborkaBarony struct {
	feud.BaseBarony
}

var BZaborka扎博尔卡 feud.Barony = &扎博尔卡ZaborkaBarony{}

func init() {
    f := BZaborka扎博尔卡.(*扎博尔卡ZaborkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaborka",
		TitleName: "扎博尔卡",
		TitleCode: "b_zaborka",
	}
}
