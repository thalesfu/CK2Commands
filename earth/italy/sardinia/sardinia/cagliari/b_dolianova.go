package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多利亚诺瓦DolianovaBarony struct {
	feud.BaseBarony
}

var BDolianova多利亚诺瓦 feud.Barony = &多利亚诺瓦DolianovaBarony{}

func init() {
	f := BDolianova多利亚诺瓦.(*多利亚诺瓦DolianovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolianova",
		TitleName: "多利亚诺瓦",
		TitleCode: "b_dolianova",
	}
}
