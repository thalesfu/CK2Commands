package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦里亚AwariyaBarony struct {
	feud.BaseBarony
}

var BAwariya阿瓦里亚 feud.Barony = &阿瓦里亚AwariyaBarony{}

func init() {
	f := BAwariya阿瓦里亚.(*阿瓦里亚AwariyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awariya",
		TitleName: "阿瓦里亚",
		TitleCode: "b_awariya",
	}
}
