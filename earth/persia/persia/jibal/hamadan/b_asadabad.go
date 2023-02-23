package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨达巴德AsadabadBarony struct {
	feud.BaseBarony
}

var BAsadabad阿萨达巴德 feud.Barony = &阿萨达巴德AsadabadBarony{}

func init() {
	f := BAsadabad阿萨达巴德.(*阿萨达巴德AsadabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asadabad",
		TitleName: "阿萨达巴德",
		TitleCode: "b_asadabad",
	}
}
