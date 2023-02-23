package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库苏亚KusuyaBarony struct {
	feud.BaseBarony
}

var BKusuya库苏亚 feud.Barony = &库苏亚KusuyaBarony{}

func init() {
	f := BKusuya库苏亚.(*库苏亚KusuyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusuya",
		TitleName: "库苏亚",
		TitleCode: "b_kusuya",
	}
}
