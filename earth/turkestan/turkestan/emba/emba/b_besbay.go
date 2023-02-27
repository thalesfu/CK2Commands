package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别斯拜BesbayBarony struct {
	feud.BaseBarony
}

var BBesbay别斯拜 feud.Barony = &别斯拜BesbayBarony{}

func init() {
    f := BBesbay别斯拜.(*别斯拜BesbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besbay",
		TitleName: "别斯拜",
		TitleCode: "b_besbay",
	}
}
