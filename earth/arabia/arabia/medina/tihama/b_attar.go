package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔尔AttarBarony struct {
	feud.BaseBarony
}

var BAttar阿塔尔 feud.Barony = &阿塔尔AttarBarony{}

func init() {
	f := BAttar阿塔尔.(*阿塔尔AttarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attar",
		TitleName: "阿塔尔",
		TitleCode: "b_attar",
	}
}
