package sharukan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔卡KharkaBarony struct {
	feud.BaseBarony
}

var BKharka哈尔卡 feud.Barony = &哈尔卡KharkaBarony{}

func init() {
    f := BKharka哈尔卡.(*哈尔卡KharkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharka",
		TitleName: "哈尔卡",
		TitleCode: "b_kharka",
	}
}
