package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌塞罗UceroBarony struct {
	feud.BaseBarony
}

var BUcero乌塞罗 feud.Barony = &乌塞罗UceroBarony{}

func init() {
	f := BUcero乌塞罗.(*乌塞罗UceroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ucero",
		TitleName: "乌塞罗",
		TitleCode: "b_ucero",
	}
}
