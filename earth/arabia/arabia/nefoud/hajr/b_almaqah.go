package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔玛达AlmaqahBarony struct {
	feud.BaseBarony
}

var BAlmaqah阿尔玛达 feud.Barony = &阿尔玛达AlmaqahBarony{}

func init() {
	f := BAlmaqah阿尔玛达.(*阿尔玛达AlmaqahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almaqah",
		TitleName: "阿尔玛达",
		TitleCode: "b_almaqah",
	}
}
