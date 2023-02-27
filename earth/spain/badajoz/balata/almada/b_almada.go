package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔马达AlmadaBarony struct {
	feud.BaseBarony
}

var BAlmada阿尔马达 feud.Barony = &阿尔马达AlmadaBarony{}

func init() {
    f := BAlmada阿尔马达.(*阿尔马达AlmadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almada",
		TitleName: "阿尔马达",
		TitleCode: "b_almada",
	}
}
