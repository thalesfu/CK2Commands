package lleida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列伊达LleidaBarony struct {
	feud.BaseBarony
}

var BLleida列伊达 feud.Barony = &列伊达LleidaBarony{}

func init() {
	f := BLleida列伊达.(*列伊达LleidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lleida",
		TitleName: "列伊达",
		TitleCode: "b_lleida",
	}
}
