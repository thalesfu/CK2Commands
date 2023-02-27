package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷沃TrevesBarony struct {
	feud.BaseBarony
}

var BTreves特雷沃 feud.Barony = &特雷沃TrevesBarony{}

func init() {
    f := BTreves特雷沃.(*特雷沃TrevesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treves",
		TitleName: "特雷沃",
		TitleCode: "b_treves",
	}
}
