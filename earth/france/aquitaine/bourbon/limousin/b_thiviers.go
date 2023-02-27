package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂维耶ThiviersBarony struct {
	feud.BaseBarony
}

var BThiviers蒂维耶 feud.Barony = &蒂维耶ThiviersBarony{}

func init() {
    f := BThiviers蒂维耶.(*蒂维耶ThiviersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thiviers",
		TitleName: "蒂维耶",
		TitleCode: "b_thiviers",
	}
}
