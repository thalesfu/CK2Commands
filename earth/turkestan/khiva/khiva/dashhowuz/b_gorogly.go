package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉奥罗格雷GoroglyBarony struct {
	feud.BaseBarony
}

var BGorogly吉奥罗格雷 feud.Barony = &吉奥罗格雷GoroglyBarony{}

func init() {
    f := BGorogly吉奥罗格雷.(*吉奥罗格雷GoroglyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorogly",
		TitleName: "吉奥罗格雷",
		TitleCode: "b_gorogly",
	}
}
