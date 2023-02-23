package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮卡廖沃PikalyovoBarony struct {
	feud.BaseBarony
}

var BPikalyovo皮卡廖沃 feud.Barony = &皮卡廖沃PikalyovoBarony{}

func init() {
	f := BPikalyovo皮卡廖沃.(*皮卡廖沃PikalyovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pikalyovo",
		TitleName: "皮卡廖沃",
		TitleCode: "b_pikalyovo",
	}
}
