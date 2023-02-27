package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 团结TuanjieBarony struct {
	feud.BaseBarony
}

var BTuanjie团结 feud.Barony = &团结TuanjieBarony{}

func init() {
    f := BTuanjie团结.(*团结TuanjieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuanjie",
		TitleName: "团结",
		TitleCode: "b_tuanjie",
	}
}
