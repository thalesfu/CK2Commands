package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马儿敢MarkamBarony struct {
	feud.BaseBarony
}

var BMarkam马儿敢 feud.Barony = &马儿敢MarkamBarony{}

func init() {
	f := BMarkam马儿敢.(*马儿敢MarkamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "markam",
		TitleName: "马儿敢",
		TitleCode: "b_markam",
	}
}
