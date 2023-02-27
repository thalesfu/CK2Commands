package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多西诺韦茨PodosinovetsBarony struct {
	feud.BaseBarony
}

var BPodosinovets波多西诺韦茨 feud.Barony = &波多西诺韦茨PodosinovetsBarony{}

func init() {
    f := BPodosinovets波多西诺韦茨.(*波多西诺韦茨PodosinovetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podosinovets",
		TitleName: "波多西诺韦茨",
		TitleCode: "b_podosinovets",
	}
}
