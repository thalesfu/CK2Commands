package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯穆伦KusmurynBarony struct {
	feud.BaseBarony
}

var BKusmuryn库斯穆伦 feud.Barony = &库斯穆伦KusmurynBarony{}

func init() {
    f := BKusmuryn库斯穆伦.(*库斯穆伦KusmurynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusmuryn",
		TitleName: "库斯穆伦",
		TitleCode: "b_kusmuryn",
	}
}
