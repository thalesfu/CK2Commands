package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛斯勒MaslBarony struct {
	feud.BaseBarony
}

var BMasl玛斯勒 feud.Barony = &玛斯勒MaslBarony{}

func init() {
    f := BMasl玛斯勒.(*玛斯勒MaslBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masl",
		TitleName: "玛斯勒",
		TitleCode: "b_masl",
	}
}
