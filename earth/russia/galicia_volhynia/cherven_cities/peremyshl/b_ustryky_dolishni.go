package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下乌斯奇基Ustryky_dolishniBarony struct {
	feud.BaseBarony
}

var BUstryky_dolishni下乌斯奇基 feud.Barony = &下乌斯奇基Ustryky_dolishniBarony{}

func init() {
    f := BUstryky_dolishni下乌斯奇基.(*下乌斯奇基Ustryky_dolishniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustryky_dolishni",
		TitleName: "下乌斯奇基",
		TitleCode: "b_ustryky_dolishni",
	}
}
