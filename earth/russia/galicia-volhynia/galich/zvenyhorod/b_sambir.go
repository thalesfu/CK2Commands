package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑博尔SambirBarony struct {
	feud.BaseBarony
}

var BSambir桑博尔 feud.Barony = &桑博尔SambirBarony{}

func init() {
    f := BSambir桑博尔.(*桑博尔SambirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sambir",
		TitleName: "桑博尔",
		TitleCode: "b_sambir",
	}
}
