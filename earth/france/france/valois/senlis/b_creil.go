package senlis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷伊CreilBarony struct {
	feud.BaseBarony
}

var BCreil克雷伊 feud.Barony = &克雷伊CreilBarony{}

func init() {
    f := BCreil克雷伊.(*克雷伊CreilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "creil",
		TitleName: "克雷伊",
		TitleCode: "b_creil",
	}
}
