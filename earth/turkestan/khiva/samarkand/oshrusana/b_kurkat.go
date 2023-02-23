package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔卡特KurkatBarony struct {
	feud.BaseBarony
}

var BKurkat古尔卡特 feud.Barony = &古尔卡特KurkatBarony{}

func init() {
	f := BKurkat古尔卡特.(*古尔卡特KurkatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurkat",
		TitleName: "古尔卡特",
		TitleCode: "b_kurkat",
	}
}
