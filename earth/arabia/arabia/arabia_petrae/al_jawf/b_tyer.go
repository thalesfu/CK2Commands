package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图耶TyerBarony struct {
	feud.BaseBarony
}

var BTyer图耶 feud.Barony = &图耶TyerBarony{}

func init() {
    f := BTyer图耶.(*图耶TyerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyer",
		TitleName: "图耶",
		TitleCode: "b_tyer",
	}
}
