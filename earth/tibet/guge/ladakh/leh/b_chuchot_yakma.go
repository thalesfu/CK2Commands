package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗秋亚玛Chuchot_yakmaBarony struct {
	feud.BaseBarony
}

var BChuchot_yakma宗秋亚玛 feud.Barony = &宗秋亚玛Chuchot_yakmaBarony{}

func init() {
    f := BChuchot_yakma宗秋亚玛.(*宗秋亚玛Chuchot_yakmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chuchot_yakma",
		TitleName: "宗秋亚玛",
		TitleCode: "b_chuchot_yakma",
	}
}
