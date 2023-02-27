package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦士AhvazBarony struct {
	feud.BaseBarony
}

var BAhvaz阿瓦士 feud.Barony = &阿瓦士AhvazBarony{}

func init() {
    f := BAhvaz阿瓦士.(*阿瓦士AhvazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahvaz",
		TitleName: "阿瓦士",
		TitleCode: "b_ahvaz",
	}
}
