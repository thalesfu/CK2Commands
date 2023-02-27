package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 慕尼黑MunchenBarony struct {
	feud.BaseBarony
}

var BMunchen慕尼黑 feud.Barony = &慕尼黑MunchenBarony{}

func init() {
    f := BMunchen慕尼黑.(*慕尼黑MunchenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munchen",
		TitleName: "慕尼黑",
		TitleCode: "b_munchen",
	}
}
