package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奎AcquiBarony struct {
	feud.BaseBarony
}

var BAcqui阿奎 feud.Barony = &阿奎AcquiBarony{}

func init() {
	f := BAcqui阿奎.(*阿奎AcquiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acqui",
		TitleName: "阿奎",
		TitleCode: "b_acqui",
	}
}
