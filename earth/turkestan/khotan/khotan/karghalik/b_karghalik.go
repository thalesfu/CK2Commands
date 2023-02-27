package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀格勒克KarghalikBarony struct {
	feud.BaseBarony
}

var BKarghalik喀格勒克 feud.Barony = &喀格勒克KarghalikBarony{}

func init() {
    f := BKarghalik喀格勒克.(*喀格勒克KarghalikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karghalik",
		TitleName: "喀格勒克",
		TitleCode: "b_karghalik",
	}
}
