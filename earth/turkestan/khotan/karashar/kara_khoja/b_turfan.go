package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吐鲁番TurfanBarony struct {
	feud.BaseBarony
}

var BTurfan吐鲁番 feud.Barony = &吐鲁番TurfanBarony{}

func init() {
    f := BTurfan吐鲁番.(*吐鲁番TurfanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turfan",
		TitleName: "吐鲁番",
		TitleCode: "b_turfan",
	}
}
