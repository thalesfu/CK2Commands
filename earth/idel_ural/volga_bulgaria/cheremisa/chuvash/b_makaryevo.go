package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马卡里耶沃MakaryevoBarony struct {
	feud.BaseBarony
}

var BMakaryevo马卡里耶沃 feud.Barony = &马卡里耶沃MakaryevoBarony{}

func init() {
    f := BMakaryevo马卡里耶沃.(*马卡里耶沃MakaryevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makaryevo",
		TitleName: "马卡里耶沃",
		TitleCode: "b_makaryevo",
	}
}
