package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔巴AlbaBarony struct {
	feud.BaseBarony
}

var BAlba阿尔巴 feud.Barony = &阿尔巴AlbaBarony{}

func init() {
	f := BAlba阿尔巴.(*阿尔巴AlbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alba",
		TitleName: "阿尔巴",
		TitleCode: "b_alba",
	}
}
