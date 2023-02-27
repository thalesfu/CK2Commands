package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔莱ArlayBarony struct {
	feud.BaseBarony
}

var BArlay阿尔莱 feud.Barony = &阿尔莱ArlayBarony{}

func init() {
    f := BArlay阿尔莱.(*阿尔莱ArlayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arlay",
		TitleName: "阿尔莱",
		TitleCode: "b_arlay",
	}
}
