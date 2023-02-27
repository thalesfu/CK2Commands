package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁塔亚奥瑟皮KrutayaosypBarony struct {
	feud.BaseBarony
}

var BKrutayaosyp克鲁塔亚奥瑟皮 feud.Barony = &克鲁塔亚奥瑟皮KrutayaosypBarony{}

func init() {
    f := BKrutayaosyp克鲁塔亚奥瑟皮.(*克鲁塔亚奥瑟皮KrutayaosypBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krutayaosyp",
		TitleName: "克鲁塔亚奥瑟皮",
		TitleCode: "b_krutayaosyp",
	}
}
