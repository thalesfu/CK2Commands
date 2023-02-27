package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克兰AucklandBarony struct {
	feud.BaseBarony
}

var BAuckland奥克兰 feud.Barony = &奥克兰AucklandBarony{}

func init() {
    f := BAuckland奥克兰.(*奥克兰AucklandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auckland",
		TitleName: "奥克兰",
		TitleCode: "b_auckland",
	}
}
