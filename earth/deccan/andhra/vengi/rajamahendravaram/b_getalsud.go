package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祇呾罗窣陀GetalsudBarony struct {
	feud.BaseBarony
}

var BGetalsud祇呾罗窣陀 feud.Barony = &祇呾罗窣陀GetalsudBarony{}

func init() {
	f := BGetalsud祇呾罗窣陀.(*祇呾罗窣陀GetalsudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "getalsud",
		TitleName: "祇呾罗窣陀",
		TitleCode: "b_getalsud",
	}
}
