package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嗢昆OrkhonBarony struct {
	feud.BaseBarony
}

var BOrkhon嗢昆 feud.Barony = &嗢昆OrkhonBarony{}

func init() {
    f := BOrkhon嗢昆.(*嗢昆OrkhonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orkhon",
		TitleName: "嗢昆",
		TitleCode: "b_orkhon",
	}
}
