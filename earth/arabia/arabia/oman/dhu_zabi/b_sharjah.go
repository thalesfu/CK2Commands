package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙迦SharjahBarony struct {
	feud.BaseBarony
}

var BSharjah沙迦 feud.Barony = &沙迦SharjahBarony{}

func init() {
    f := BSharjah沙迦.(*沙迦SharjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharjah",
		TitleName: "沙迦",
		TitleCode: "b_sharjah",
	}
}
