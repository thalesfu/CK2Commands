package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂西塞姆EnsisheimBarony struct {
	feud.BaseBarony
}

var BEnsisheim昂西塞姆 feud.Barony = &昂西塞姆EnsisheimBarony{}

func init() {
    f := BEnsisheim昂西塞姆.(*昂西塞姆EnsisheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ensisheim",
		TitleName: "昂西塞姆",
		TitleCode: "b_ensisheim",
	}
}
