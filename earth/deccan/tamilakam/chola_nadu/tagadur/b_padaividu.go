package tagadur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多毗都PadaividuBarony struct {
	feud.BaseBarony
}

var BPadaividu波多毗都 feud.Barony = &波多毗都PadaividuBarony{}

func init() {
    f := BPadaividu波多毗都.(*波多毗都PadaividuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "padaividu",
		TitleName: "波多毗都",
		TitleCode: "b_padaividu",
	}
}
