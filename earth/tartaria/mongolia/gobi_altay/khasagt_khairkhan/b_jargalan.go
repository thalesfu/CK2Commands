package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔嘎朗JargalanBarony struct {
	feud.BaseBarony
}

var BJargalan扎尔嘎朗 feud.Barony = &扎尔嘎朗JargalanBarony{}

func init() {
    f := BJargalan扎尔嘎朗.(*扎尔嘎朗JargalanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jargalan",
		TitleName: "扎尔嘎朗",
		TitleCode: "b_jargalan",
	}
}
