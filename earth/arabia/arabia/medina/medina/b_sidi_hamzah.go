package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪哈穆扎Sidi_hamzahBarony struct {
	feud.BaseBarony
}

var BSidi_hamzah西迪哈穆扎 feud.Barony = &西迪哈穆扎Sidi_hamzahBarony{}

func init() {
    f := BSidi_hamzah西迪哈穆扎.(*西迪哈穆扎Sidi_hamzahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidi_hamzah",
		TitleName: "西迪哈穆扎",
		TitleCode: "b_sidi_hamzah",
	}
}
