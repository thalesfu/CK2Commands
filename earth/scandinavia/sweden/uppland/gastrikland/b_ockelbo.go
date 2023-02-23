package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克尔布OckelboBarony struct {
	feud.BaseBarony
}

var BOckelbo奥克尔布 feud.Barony = &奥克尔布OckelboBarony{}

func init() {
	f := BOckelbo奥克尔布.(*奥克尔布OckelboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ockelbo",
		TitleName: "奥克尔布",
		TitleCode: "b_ockelbo",
	}
}
