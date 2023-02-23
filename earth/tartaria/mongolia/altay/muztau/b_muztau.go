package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木斯岛MuztauBarony struct {
	feud.BaseBarony
}

var BMuztau木斯岛 feud.Barony = &木斯岛MuztauBarony{}

func init() {
	f := BMuztau木斯岛.(*木斯岛MuztauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muztau",
		TitleName: "木斯岛",
		TitleCode: "b_muztau",
	}
}
