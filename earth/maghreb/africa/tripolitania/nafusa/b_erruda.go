package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃鲁达ErrudaBarony struct {
	feud.BaseBarony
}

var BErruda埃鲁达 feud.Barony = &埃鲁达ErrudaBarony{}

func init() {
	f := BErruda埃鲁达.(*埃鲁达ErrudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erruda",
		TitleName: "埃鲁达",
		TitleCode: "b_erruda",
	}
}
