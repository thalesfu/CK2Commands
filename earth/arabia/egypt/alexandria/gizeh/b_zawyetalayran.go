package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎维耶奥亚因ZawyetalayranBarony struct {
	feud.BaseBarony
}

var BZawyetalayran扎维耶奥亚因 feud.Barony = &扎维耶奥亚因ZawyetalayranBarony{}

func init() {
	f := BZawyetalayran扎维耶奥亚因.(*扎维耶奥亚因ZawyetalayranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zawyetalayran",
		TitleName: "扎维耶奥亚因",
		TitleCode: "b_zawyetalayran",
	}
}
