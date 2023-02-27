package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚东YatungBarony struct {
	feud.BaseBarony
}

var BYatung亚东 feud.Barony = &亚东YatungBarony{}

func init() {
    f := BYatung亚东.(*亚东YatungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yatung",
		TitleName: "亚东",
		TitleCode: "b_yatung",
	}
}
