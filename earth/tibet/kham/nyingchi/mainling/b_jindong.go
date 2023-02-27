package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金东JindongBarony struct {
	feud.BaseBarony
}

var BJindong金东 feud.Barony = &金东JindongBarony{}

func init() {
    f := BJindong金东.(*金东JindongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jindong",
		TitleName: "金东",
		TitleCode: "b_jindong",
	}
}
