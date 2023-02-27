package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达纳AdanaBarony struct {
	feud.BaseBarony
}

var BAdana阿达纳 feud.Barony = &阿达纳AdanaBarony{}

func init() {
    f := BAdana阿达纳.(*阿达纳AdanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adana",
		TitleName: "阿达纳",
		TitleCode: "b_adana",
	}
}
