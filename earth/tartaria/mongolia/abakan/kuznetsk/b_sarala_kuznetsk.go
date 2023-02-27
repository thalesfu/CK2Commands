package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉拉Sarala_kuznetskBarony struct {
	feud.BaseBarony
}

var BSarala_kuznetsk萨拉拉 feud.Barony = &萨拉拉Sarala_kuznetskBarony{}

func init() {
    f := BSarala_kuznetsk萨拉拉.(*萨拉拉Sarala_kuznetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarala_kuznetsk",
		TitleName: "萨拉拉",
		TitleCode: "b_sarala_kuznetsk",
	}
}
