package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷兹尼克DreznikBarony struct {
	feud.BaseBarony
}

var BDreznik德雷兹尼克 feud.Barony = &德雷兹尼克DreznikBarony{}

func init() {
    f := BDreznik德雷兹尼克.(*德雷兹尼克DreznikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dreznik",
		TitleName: "德雷兹尼克",
		TitleCode: "b_dreznik",
	}
}
