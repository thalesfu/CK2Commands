package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿曼克孜利特AmankyzylitBarony struct {
	feud.BaseBarony
}

var BAmankyzylit阿曼克孜利特 feud.Barony = &阿曼克孜利特AmankyzylitBarony{}

func init() {
    f := BAmankyzylit阿曼克孜利特.(*阿曼克孜利特AmankyzylitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amankyzylit",
		TitleName: "阿曼克孜利特",
		TitleCode: "b_amankyzylit",
	}
}
