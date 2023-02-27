package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦提AwatBarony struct {
	feud.BaseBarony
}

var BAwat阿瓦提 feud.Barony = &阿瓦提AwatBarony{}

func init() {
    f := BAwat阿瓦提.(*阿瓦提AwatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awat",
		TitleName: "阿瓦提",
		TitleCode: "b_awat",
	}
}
