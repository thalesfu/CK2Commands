package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔扣克SelkirkBarony struct {
	feud.BaseBarony
}

var BSelkirk塞尔扣克 feud.Barony = &塞尔扣克SelkirkBarony{}

func init() {
	f := BSelkirk塞尔扣克.(*塞尔扣克SelkirkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selkirk",
		TitleName: "塞尔扣克",
		TitleCode: "b_selkirk",
	}
}
