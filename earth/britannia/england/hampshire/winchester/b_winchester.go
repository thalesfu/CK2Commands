package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温彻斯特WinchesterBarony struct {
	feud.BaseBarony
}

var BWinchester温彻斯特 feud.Barony = &温彻斯特WinchesterBarony{}

func init() {
	f := BWinchester温彻斯特.(*温彻斯特WinchesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "winchester",
		TitleName: "温彻斯特",
		TitleCode: "b_winchester",
	}
}
