package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔赫西拉斯AlgericasBarony struct {
	feud.BaseBarony
}

var BAlgericas阿尔赫西拉斯 feud.Barony = &阿尔赫西拉斯AlgericasBarony{}

func init() {
	f := BAlgericas阿尔赫西拉斯.(*阿尔赫西拉斯AlgericasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "algericas",
		TitleName: "阿尔赫西拉斯",
		TitleCode: "b_algericas",
	}
}
