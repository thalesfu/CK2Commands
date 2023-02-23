package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷林HlerinBarony struct {
	feud.BaseBarony
}

var BHlerin赫雷林 feud.Barony = &赫雷林HlerinBarony{}

func init() {
	f := BHlerin赫雷林.(*赫雷林HlerinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hlerin",
		TitleName: "赫雷林",
		TitleCode: "b_hlerin",
	}
}
