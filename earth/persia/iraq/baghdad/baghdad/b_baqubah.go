package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴古拜BaqubahBarony struct {
	feud.BaseBarony
}

var BBaqubah巴古拜 feud.Barony = &巴古拜BaqubahBarony{}

func init() {
	f := BBaqubah巴古拜.(*巴古拜BaqubahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baqubah",
		TitleName: "巴古拜",
		TitleCode: "b_baqubah",
	}
}
