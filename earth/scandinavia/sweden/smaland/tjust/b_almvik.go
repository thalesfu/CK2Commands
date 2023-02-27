package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔姆维克AlmvikBarony struct {
	feud.BaseBarony
}

var BAlmvik阿尔姆维克 feud.Barony = &阿尔姆维克AlmvikBarony{}

func init() {
    f := BAlmvik阿尔姆维克.(*阿尔姆维克AlmvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almvik",
		TitleName: "阿尔姆维克",
		TitleCode: "b_almvik",
	}
}
