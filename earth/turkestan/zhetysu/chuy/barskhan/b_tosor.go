package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托索尔TosorBarony struct {
	feud.BaseBarony
}

var BTosor托索尔 feud.Barony = &托索尔TosorBarony{}

func init() {
    f := BTosor托索尔.(*托索尔TosorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tosor",
		TitleName: "托索尔",
		TitleCode: "b_tosor",
	}
}
