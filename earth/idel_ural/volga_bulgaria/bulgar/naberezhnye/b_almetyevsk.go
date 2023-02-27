package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔梅季耶夫斯克AlmetyevskBarony struct {
	feud.BaseBarony
}

var BAlmetyevsk阿尔梅季耶夫斯克 feud.Barony = &阿尔梅季耶夫斯克AlmetyevskBarony{}

func init() {
    f := BAlmetyevsk阿尔梅季耶夫斯克.(*阿尔梅季耶夫斯克AlmetyevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almetyevsk",
		TitleName: "阿尔梅季耶夫斯克",
		TitleCode: "b_almetyevsk",
	}
}
