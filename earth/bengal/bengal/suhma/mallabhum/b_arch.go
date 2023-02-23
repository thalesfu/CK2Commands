package mallabhum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿遮ArchBarony struct {
	feud.BaseBarony
}

var BArch阿遮 feud.Barony = &阿遮ArchBarony{}

func init() {
	f := BArch阿遮.(*阿遮ArchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arch",
		TitleName: "阿遮",
		TitleCode: "b_arch",
	}
}
