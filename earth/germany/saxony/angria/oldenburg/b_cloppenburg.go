package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛彭堡CloppenburgBarony struct {
	feud.BaseBarony
}

var BCloppenburg克洛彭堡 feud.Barony = &克洛彭堡CloppenburgBarony{}

func init() {
	f := BCloppenburg克洛彭堡.(*克洛彭堡CloppenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cloppenburg",
		TitleName: "克洛彭堡",
		TitleCode: "b_cloppenburg",
	}
}
