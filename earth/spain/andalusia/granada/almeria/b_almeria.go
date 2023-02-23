package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔梅里亚AlmeriaBarony struct {
	feud.BaseBarony
}

var BAlmeria阿尔梅里亚 feud.Barony = &阿尔梅里亚AlmeriaBarony{}

func init() {
	f := BAlmeria阿尔梅里亚.(*阿尔梅里亚AlmeriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almeria",
		TitleName: "阿尔梅里亚",
		TitleCode: "b_almeria",
	}
}
