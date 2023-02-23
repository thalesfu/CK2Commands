package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基普加伊KipgajBarony struct {
	feud.BaseBarony
}

var BKipgaj基普加伊 feud.Barony = &基普加伊KipgajBarony{}

func init() {
	f := BKipgaj基普加伊.(*基普加伊KipgajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kipgaj",
		TitleName: "基普加伊",
		TitleCode: "b_kipgaj",
	}
}
