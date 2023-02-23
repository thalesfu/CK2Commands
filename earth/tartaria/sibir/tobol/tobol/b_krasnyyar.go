package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯内_亚尔KrasnyyarBarony struct {
	feud.BaseBarony
}

var BKrasnyyar克拉斯内_亚尔 feud.Barony = &克拉斯内_亚尔KrasnyyarBarony{}

func init() {
	f := BKrasnyyar克拉斯内_亚尔.(*克拉斯内_亚尔KrasnyyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnyyar",
		TitleName: "克拉斯内_亚尔",
		TitleCode: "b_krasnyyar",
	}
}
