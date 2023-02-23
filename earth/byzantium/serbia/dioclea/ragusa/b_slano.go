package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯拉诺SlanoBarony struct {
	feud.BaseBarony
}

var BSlano斯拉诺 feud.Barony = &斯拉诺SlanoBarony{}

func init() {
	f := BSlano斯拉诺.(*斯拉诺SlanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slano",
		TitleName: "斯拉诺",
		TitleCode: "b_slano",
	}
}
