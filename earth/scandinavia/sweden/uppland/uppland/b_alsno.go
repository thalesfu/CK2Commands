package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯内AlsnoBarony struct {
	feud.BaseBarony
}

var BAlsno阿尔斯内 feud.Barony = &阿尔斯内AlsnoBarony{}

func init() {
	f := BAlsno阿尔斯内.(*阿尔斯内AlsnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsno",
		TitleName: "阿尔斯内",
		TitleCode: "b_alsno",
	}
}
