package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓弗姆林DunfermlineBarony struct {
	feud.BaseBarony
}

var BDunfermline邓弗姆林 feud.Barony = &邓弗姆林DunfermlineBarony{}

func init() {
	f := BDunfermline邓弗姆林.(*邓弗姆林DunfermlineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunfermline",
		TitleName: "邓弗姆林",
		TitleCode: "b_dunfermline",
	}
}
