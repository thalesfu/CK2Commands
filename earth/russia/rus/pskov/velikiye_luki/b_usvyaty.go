package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯维亚特UsvyatyBarony struct {
	feud.BaseBarony
}

var BUsvyaty乌斯维亚特 feud.Barony = &乌斯维亚特UsvyatyBarony{}

func init() {
    f := BUsvyaty乌斯维亚特.(*乌斯维亚特UsvyatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usvyaty",
		TitleName: "乌斯维亚特",
		TitleCode: "b_usvyaty",
	}
}
