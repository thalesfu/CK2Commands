package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿勒贾尔AlgiarBarony struct {
	feud.BaseBarony
}

var BAlgiar阿勒贾尔 feud.Barony = &阿勒贾尔AlgiarBarony{}

func init() {
	f := BAlgiar阿勒贾尔.(*阿勒贾尔AlgiarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algiar",
		TitleName: "阿勒贾尔",
		TitleCode: "b_algiar",
	}
}
