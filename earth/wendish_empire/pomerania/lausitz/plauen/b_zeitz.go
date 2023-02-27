package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蔡茨ZeitzBarony struct {
	feud.BaseBarony
}

var BZeitz蔡茨 feud.Barony = &蔡茨ZeitzBarony{}

func init() {
    f := BZeitz蔡茨.(*蔡茨ZeitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeitz",
		TitleName: "蔡茨",
		TitleCode: "b_zeitz",
	}
}
