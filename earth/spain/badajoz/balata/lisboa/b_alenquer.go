package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伦克尔AlenquerBarony struct {
	feud.BaseBarony
}

var BAlenquer阿伦克尔 feud.Barony = &阿伦克尔AlenquerBarony{}

func init() {
    f := BAlenquer阿伦克尔.(*阿伦克尔AlenquerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alenquer",
		TitleName: "阿伦克尔",
		TitleCode: "b_alenquer",
	}
}
