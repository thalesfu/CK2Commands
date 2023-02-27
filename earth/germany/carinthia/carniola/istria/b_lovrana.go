package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛夫拉纳LovranaBarony struct {
	feud.BaseBarony
}

var BLovrana洛夫拉纳 feud.Barony = &洛夫拉纳LovranaBarony{}

func init() {
    f := BLovrana洛夫拉纳.(*洛夫拉纳LovranaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lovrana",
		TitleName: "洛夫拉纳",
		TitleCode: "b_lovrana",
	}
}
