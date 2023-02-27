package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔尔瑙CsornaBarony struct {
	feud.BaseBarony
}

var BCsorna乔尔瑙 feud.Barony = &乔尔瑙CsornaBarony{}

func init() {
    f := BCsorna乔尔瑙.(*乔尔瑙CsornaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "csorna",
		TitleName: "乔尔瑙",
		TitleCode: "b_csorna",
	}
}
