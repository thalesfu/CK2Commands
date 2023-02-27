package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 答丽河DaliheBarony struct {
	feud.BaseBarony
}

var BDalihe答丽河 feud.Barony = &答丽河DaliheBarony{}

func init() {
    f := BDalihe答丽河.(*答丽河DaliheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalihe",
		TitleName: "答丽河",
		TitleCode: "b_dalihe",
	}
}
