package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔布开克AlburquerqueBarony struct {
	feud.BaseBarony
}

var BAlburquerque阿尔布开克 feud.Barony = &阿尔布开克AlburquerqueBarony{}

func init() {
	f := BAlburquerque阿尔布开克.(*阿尔布开克AlburquerqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alburquerque",
		TitleName: "阿尔布开克",
		TitleCode: "b_alburquerque",
	}
}
