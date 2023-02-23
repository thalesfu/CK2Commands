package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔曼萨AlmansaBarony struct {
	feud.BaseBarony
}

var BAlmansa阿尔曼萨 feud.Barony = &阿尔曼萨AlmansaBarony{}

func init() {
	f := BAlmansa阿尔曼萨.(*阿尔曼萨AlmansaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almansa",
		TitleName: "阿尔曼萨",
		TitleCode: "b_almansa",
	}
}
