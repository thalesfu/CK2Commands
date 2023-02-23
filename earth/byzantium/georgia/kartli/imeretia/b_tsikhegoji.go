package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐赫戈吉TsikhegojiBarony struct {
	feud.BaseBarony
}

var BTsikhegoji齐赫戈吉 feud.Barony = &齐赫戈吉TsikhegojiBarony{}

func init() {
	f := BTsikhegoji齐赫戈吉.(*齐赫戈吉TsikhegojiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsikhegoji",
		TitleName: "齐赫戈吉",
		TitleCode: "b_tsikhegoji",
	}
}
