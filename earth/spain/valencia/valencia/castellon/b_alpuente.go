package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔普恩特AlpuenteBarony struct {
	feud.BaseBarony
}

var BAlpuente阿尔普恩特 feud.Barony = &阿尔普恩特AlpuenteBarony{}

func init() {
    f := BAlpuente阿尔普恩特.(*阿尔普恩特AlpuenteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alpuente",
		TitleName: "阿尔普恩特",
		TitleCode: "b_alpuente",
	}
}
