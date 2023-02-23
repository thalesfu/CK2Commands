package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 冲哈ChhonhupBarony struct {
	feud.BaseBarony
}

var BChhonhup冲哈 feud.Barony = &冲哈ChhonhupBarony{}

func init() {
	f := BChhonhup冲哈.(*冲哈ChhonhupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chhonhup",
		TitleName: "冲哈",
		TitleCode: "b_chhonhup",
	}
}
