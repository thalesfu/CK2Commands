package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克雷AnchoryBarony struct {
	feud.BaseBarony
}

var BAnchory阿克雷 feud.Barony = &阿克雷AnchoryBarony{}

func init() {
	f := BAnchory阿克雷.(*阿克雷AnchoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anchory",
		TitleName: "阿克雷",
		TitleCode: "b_anchory",
	}
}
