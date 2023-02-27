package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔盖罗AlgheroBarony struct {
	feud.BaseBarony
}

var BAlghero阿尔盖罗 feud.Barony = &阿尔盖罗AlgheroBarony{}

func init() {
    f := BAlghero阿尔盖罗.(*阿尔盖罗AlgheroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alghero",
		TitleName: "阿尔盖罗",
		TitleCode: "b_alghero",
	}
}
