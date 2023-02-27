package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎阿尔赫JaarahBarony struct {
	feud.BaseBarony
}

var BJaarah扎阿尔赫 feud.Barony = &扎阿尔赫JaarahBarony{}

func init() {
    f := BJaarah扎阿尔赫.(*扎阿尔赫JaarahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaarah",
		TitleName: "扎阿尔赫",
		TitleCode: "b_jaarah",
	}
}
