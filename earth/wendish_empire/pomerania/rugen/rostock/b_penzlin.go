package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彭茨林PenzlinBarony struct {
	feud.BaseBarony
}

var BPenzlin彭茨林 feud.Barony = &彭茨林PenzlinBarony{}

func init() {
    f := BPenzlin彭茨林.(*彭茨林PenzlinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penzlin",
		TitleName: "彭茨林",
		TitleCode: "b_penzlin",
	}
}
