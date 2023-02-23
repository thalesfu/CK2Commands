package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛基雷克LocquirecBarony struct {
	feud.BaseBarony
}

var BLocquirec洛基雷克 feud.Barony = &洛基雷克LocquirecBarony{}

func init() {
	f := BLocquirec洛基雷克.(*洛基雷克LocquirecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "locquirec",
		TitleName: "洛基雷克",
		TitleCode: "b_locquirec",
	}
}
