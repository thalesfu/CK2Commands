package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维耶VihiersBarony struct {
	feud.BaseBarony
}

var BVihiers维耶 feud.Barony = &维耶VihiersBarony{}

func init() {
    f := BVihiers维耶.(*维耶VihiersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vihiers",
		TitleName: "维耶",
		TitleCode: "b_vihiers",
	}
}
