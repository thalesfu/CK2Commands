package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟堡CherbourgBarony struct {
	feud.BaseBarony
}

var BCherbourg瑟堡 feud.Barony = &瑟堡CherbourgBarony{}

func init() {
	f := BCherbourg瑟堡.(*瑟堡CherbourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherbourg",
		TitleName: "瑟堡",
		TitleCode: "b_cherbourg",
	}
}
