package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲登QudengBarony struct {
	feud.BaseBarony
}

var BQudeng曲登 feud.Barony = &曲登QudengBarony{}

func init() {
	f := BQudeng曲登.(*曲登QudengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qudeng",
		TitleName: "曲登",
		TitleCode: "b_qudeng",
	}
}
