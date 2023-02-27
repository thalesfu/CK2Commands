package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈里安德鲁斯MyriandrosBarony struct {
	feud.BaseBarony
}

var BMyriandros迈里安德鲁斯 feud.Barony = &迈里安德鲁斯MyriandrosBarony{}

func init() {
    f := BMyriandros迈里安德鲁斯.(*迈里安德鲁斯MyriandrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myriandros",
		TitleName: "迈里安德鲁斯",
		TitleCode: "b_myriandros",
	}
}
