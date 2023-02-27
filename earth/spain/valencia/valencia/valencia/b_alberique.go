package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔韦里克AlberiqueBarony struct {
	feud.BaseBarony
}

var BAlberique阿尔韦里克 feud.Barony = &阿尔韦里克AlberiqueBarony{}

func init() {
    f := BAlberique阿尔韦里克.(*阿尔韦里克AlberiqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alberique",
		TitleName: "阿尔韦里克",
		TitleCode: "b_alberique",
	}
}
