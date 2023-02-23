package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔塔托内CurtatoneBarony struct {
	feud.BaseBarony
}

var BCurtatone库尔塔托内 feud.Barony = &库尔塔托内CurtatoneBarony{}

func init() {
	f := BCurtatone库尔塔托内.(*库尔塔托内CurtatoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "curtatone",
		TitleName: "库尔塔托内",
		TitleCode: "b_curtatone",
	}
}
