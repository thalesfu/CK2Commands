package forcalquier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡斯特龙SisteronBarony struct {
	feud.BaseBarony
}

var BSisteron锡斯特龙 feud.Barony = &锡斯特龙SisteronBarony{}

func init() {
    f := BSisteron锡斯特龙.(*锡斯特龙SisteronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sisteron",
		TitleName: "锡斯特龙",
		TitleCode: "b_sisteron",
	}
}
