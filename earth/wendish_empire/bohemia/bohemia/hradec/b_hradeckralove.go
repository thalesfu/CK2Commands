package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫拉德茨克拉洛韦HradeckraloveBarony struct {
	feud.BaseBarony
}

var BHradeckralove赫拉德茨克拉洛韦 feud.Barony = &赫拉德茨克拉洛韦HradeckraloveBarony{}

func init() {
    f := BHradeckralove赫拉德茨克拉洛韦.(*赫拉德茨克拉洛韦HradeckraloveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hradeckralove",
		TitleName: "赫拉德茨克拉洛韦",
		TitleCode: "b_hradeckralove",
	}
}
