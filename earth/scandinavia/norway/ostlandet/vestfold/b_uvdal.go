package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌夫达UvdalBarony struct {
	feud.BaseBarony
}

var BUvdal乌夫达 feud.Barony = &乌夫达UvdalBarony{}

func init() {
	f := BUvdal乌夫达.(*乌夫达UvdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uvdal",
		TitleName: "乌夫达",
		TitleCode: "b_uvdal",
	}
}
