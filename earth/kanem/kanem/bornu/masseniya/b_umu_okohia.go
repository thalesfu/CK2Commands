package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌穆奥科希亚Umu_okohiaBarony struct {
	feud.BaseBarony
}

var BUmu_okohia乌穆奥科希亚 feud.Barony = &乌穆奥科希亚Umu_okohiaBarony{}

func init() {
    f := BUmu_okohia乌穆奥科希亚.(*乌穆奥科希亚Umu_okohiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umu_okohia",
		TitleName: "乌穆奥科希亚",
		TitleCode: "b_umu_okohia",
	}
}
