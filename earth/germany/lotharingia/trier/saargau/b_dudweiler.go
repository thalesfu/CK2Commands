package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜德韦勒DudweilerBarony struct {
	feud.BaseBarony
}

var BDudweiler杜德韦勒 feud.Barony = &杜德韦勒DudweilerBarony{}

func init() {
    f := BDudweiler杜德韦勒.(*杜德韦勒DudweilerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dudweiler",
		TitleName: "杜德韦勒",
		TitleCode: "b_dudweiler",
	}
}
