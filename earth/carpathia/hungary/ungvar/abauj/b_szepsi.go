package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞普希SzepsiBarony struct {
	feud.BaseBarony
}

var BSzepsi塞普希 feud.Barony = &塞普希SzepsiBarony{}

func init() {
    f := BSzepsi塞普希.(*塞普希SzepsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szepsi",
		TitleName: "塞普希",
		TitleCode: "b_szepsi",
	}
}
