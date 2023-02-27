package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉布里LabritBarony struct {
	feud.BaseBarony
}

var BLabrit拉布里 feud.Barony = &拉布里LabritBarony{}

func init() {
    f := BLabrit拉布里.(*拉布里LabritBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labrit",
		TitleName: "拉布里",
		TitleCode: "b_labrit",
	}
}
