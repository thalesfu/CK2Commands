package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库夫拉KufraBarony struct {
	feud.BaseBarony
}

var BKufra库夫拉 feud.Barony = &库夫拉KufraBarony{}

func init() {
    f := BKufra库夫拉.(*库夫拉KufraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kufra",
		TitleName: "库夫拉",
		TitleCode: "b_kufra",
	}
}
