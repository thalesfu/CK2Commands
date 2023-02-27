package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔嘎ShargaBarony struct {
	feud.BaseBarony
}

var BSharga沙尔嘎 feud.Barony = &沙尔嘎ShargaBarony{}

func init() {
    f := BSharga沙尔嘎.(*沙尔嘎ShargaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharga",
		TitleName: "沙尔嘎",
		TitleCode: "b_sharga",
	}
}
