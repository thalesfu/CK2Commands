package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 频阇慕罗VinjamoorBarony struct {
	feud.BaseBarony
}

var BVinjamoor频阇慕罗 feud.Barony = &频阇慕罗VinjamoorBarony{}

func init() {
	f := BVinjamoor频阇慕罗.(*频阇慕罗VinjamoorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinjamoor",
		TitleName: "频阇慕罗",
		TitleCode: "b_vinjamoor",
	}
}
