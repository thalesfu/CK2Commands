package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉拉SalalahBarony struct {
	feud.BaseBarony
}

var BSalalah萨拉拉 feud.Barony = &萨拉拉SalalahBarony{}

func init() {
    f := BSalalah萨拉拉.(*萨拉拉SalalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salalah",
		TitleName: "萨拉拉",
		TitleCode: "b_salalah",
	}
}
