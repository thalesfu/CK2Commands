package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温布里亚蒂科UmbriaticoBarony struct {
	feud.BaseBarony
}

var BUmbriatico温布里亚蒂科 feud.Barony = &温布里亚蒂科UmbriaticoBarony{}

func init() {
	f := BUmbriatico温布里亚蒂科.(*温布里亚蒂科UmbriaticoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umbriatico",
		TitleName: "温布里亚蒂科",
		TitleCode: "b_umbriatico",
	}
}
