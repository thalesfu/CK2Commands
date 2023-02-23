package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞迦审IshkashimBarony struct {
	feud.BaseBarony
}

var BIshkashim塞迦审 feud.Barony = &塞迦审IshkashimBarony{}

func init() {
	f := BIshkashim塞迦审.(*塞迦审IshkashimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishkashim",
		TitleName: "塞迦审",
		TitleCode: "b_ishkashim",
	}
}
