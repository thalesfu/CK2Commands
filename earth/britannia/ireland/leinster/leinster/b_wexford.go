package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦克斯福德WexfordBarony struct {
	feud.BaseBarony
}

var BWexford韦克斯福德 feud.Barony = &韦克斯福德WexfordBarony{}

func init() {
	f := BWexford韦克斯福德.(*韦克斯福德WexfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wexford",
		TitleName: "韦克斯福德",
		TitleCode: "b_wexford",
	}
}
