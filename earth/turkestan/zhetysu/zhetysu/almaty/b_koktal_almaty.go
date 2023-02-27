package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克塔尔Koktal_almatyBarony struct {
	feud.BaseBarony
}

var BKoktal_almaty科克塔尔 feud.Barony = &科克塔尔Koktal_almatyBarony{}

func init() {
    f := BKoktal_almaty科克塔尔.(*科克塔尔Koktal_almatyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koktal_almaty",
		TitleName: "科克塔尔",
		TitleCode: "b_koktal_almaty",
	}
}
