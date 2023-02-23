package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基辅KievBarony struct {
	feud.BaseBarony
}

var BKiev基辅 feud.Barony = &基辅KievBarony{}

func init() {
	f := BKiev基辅.(*基辅KievBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiev",
		TitleName: "基辅",
		TitleCode: "b_kiev",
	}
}
