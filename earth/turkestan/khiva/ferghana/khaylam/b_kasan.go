package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 渴塞KasanBarony struct {
	feud.BaseBarony
}

var BKasan渴塞 feud.Barony = &渴塞KasanBarony{}

func init() {
    f := BKasan渴塞.(*渴塞KasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasan",
		TitleName: "渴塞",
		TitleCode: "b_kasan",
	}
}
