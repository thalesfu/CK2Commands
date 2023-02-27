package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡斯特HusztBarony struct {
	feud.BaseBarony
}

var BHuszt胡斯特 feud.Barony = &胡斯特HusztBarony{}

func init() {
    f := BHuszt胡斯特.(*胡斯特HusztBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huszt",
		TitleName: "胡斯特",
		TitleCode: "b_huszt",
	}
}
