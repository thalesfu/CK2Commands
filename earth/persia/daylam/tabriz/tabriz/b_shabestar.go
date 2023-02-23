package tabriz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙贝斯塔尔ShabestarBarony struct {
	feud.BaseBarony
}

var BShabestar沙贝斯塔尔 feud.Barony = &沙贝斯塔尔ShabestarBarony{}

func init() {
	f := BShabestar沙贝斯塔尔.(*沙贝斯塔尔ShabestarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shabestar",
		TitleName: "沙贝斯塔尔",
		TitleCode: "b_shabestar",
	}
}
