package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 铁卜加TiebujiaBarony struct {
	feud.BaseBarony
}

var BTiebujia铁卜加 feud.Barony = &铁卜加TiebujiaBarony{}

func init() {
    f := BTiebujia铁卜加.(*铁卜加TiebujiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiebujia",
		TitleName: "铁卜加",
		TitleCode: "b_tiebujia",
	}
}
