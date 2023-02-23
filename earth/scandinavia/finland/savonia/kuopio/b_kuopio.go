package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库奥皮奥KuopioBarony struct {
	feud.BaseBarony
}

var BKuopio库奥皮奥 feud.Barony = &库奥皮奥KuopioBarony{}

func init() {
	f := BKuopio库奥皮奥.(*库奥皮奥KuopioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuopio",
		TitleName: "库奥皮奥",
		TitleCode: "b_kuopio",
	}
}
