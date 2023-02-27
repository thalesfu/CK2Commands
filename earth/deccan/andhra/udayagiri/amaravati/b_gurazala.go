package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古罗扎拉GurazalaBarony struct {
	feud.BaseBarony
}

var BGurazala古罗扎拉 feud.Barony = &古罗扎拉GurazalaBarony{}

func init() {
    f := BGurazala古罗扎拉.(*古罗扎拉GurazalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurazala",
		TitleName: "古罗扎拉",
		TitleCode: "b_gurazala",
	}
}
