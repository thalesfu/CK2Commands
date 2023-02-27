package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦萨AwasaBarony struct {
	feud.BaseBarony
}

var BAwasa阿瓦萨 feud.Barony = &阿瓦萨AwasaBarony{}

func init() {
    f := BAwasa阿瓦萨.(*阿瓦萨AwasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awasa",
		TitleName: "阿瓦萨",
		TitleCode: "b_awasa",
	}
}
