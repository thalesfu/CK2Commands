package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特奈MontenayBarony struct {
	feud.BaseBarony
}

var BMontenay蒙特奈 feud.Barony = &蒙特奈MontenayBarony{}

func init() {
    f := BMontenay蒙特奈.(*蒙特奈MontenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montenay",
		TitleName: "蒙特奈",
		TitleCode: "b_montenay",
	}
}
