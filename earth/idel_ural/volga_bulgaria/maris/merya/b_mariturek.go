package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里_图列克MariturekBarony struct {
	feud.BaseBarony
}

var BMariturek马里_图列克 feud.Barony = &马里_图列克MariturekBarony{}

func init() {
    f := BMariturek马里_图列克.(*马里_图列克MariturekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mariturek",
		TitleName: "马里_图列克",
		TitleCode: "b_mariturek",
	}
}
