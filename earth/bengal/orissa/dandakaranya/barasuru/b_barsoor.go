package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柏索尔BarsoorBarony struct {
	feud.BaseBarony
}

var BBarsoor柏索尔 feud.Barony = &柏索尔BarsoorBarony{}

func init() {
	f := BBarsoor柏索尔.(*柏索尔BarsoorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barsoor",
		TitleName: "柏索尔",
		TitleCode: "b_barsoor",
	}
}
