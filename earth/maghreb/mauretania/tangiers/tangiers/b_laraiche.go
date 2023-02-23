package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉腊什LaraicheBarony struct {
	feud.BaseBarony
}

var BLaraiche拉腊什 feud.Barony = &拉腊什LaraicheBarony{}

func init() {
	f := BLaraiche拉腊什.(*拉腊什LaraicheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laraiche",
		TitleName: "拉腊什",
		TitleCode: "b_laraiche",
	}
}
