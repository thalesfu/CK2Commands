package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆韦利赫MuwaylihBarony struct {
	feud.BaseBarony
}

var BMuwaylih穆韦利赫 feud.Barony = &穆韦利赫MuwaylihBarony{}

func init() {
    f := BMuwaylih穆韦利赫.(*穆韦利赫MuwaylihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muwaylih",
		TitleName: "穆韦利赫",
		TitleCode: "b_muwaylih",
	}
}
