package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海法尔AlhaifarBarony struct {
	feud.BaseBarony
}

var BAlhaifar海法尔 feud.Barony = &海法尔AlhaifarBarony{}

func init() {
	f := BAlhaifar海法尔.(*海法尔AlhaifarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhaifar",
		TitleName: "海法尔",
		TitleCode: "b_alhaifar",
	}
}
