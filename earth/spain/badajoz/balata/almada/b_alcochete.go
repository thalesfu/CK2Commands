package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科谢泰AlcocheteBarony struct {
	feud.BaseBarony
}

var BAlcochete阿尔科谢泰 feud.Barony = &阿尔科谢泰AlcocheteBarony{}

func init() {
	f := BAlcochete阿尔科谢泰.(*阿尔科谢泰AlcocheteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcochete",
		TitleName: "阿尔科谢泰",
		TitleCode: "b_alcochete",
	}
}
