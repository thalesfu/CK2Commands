package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴伦Baruun_saikhany_nuruuBarony struct {
	feud.BaseBarony
}

var BBaruun_saikhany_nuruu巴伦 feud.Barony = &巴伦Baruun_saikhany_nuruuBarony{}

func init() {
    f := BBaruun_saikhany_nuruu巴伦.(*巴伦Baruun_saikhany_nuruuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baruun_saikhany_nuruu",
		TitleName: "巴伦",
		TitleCode: "b_baruun_saikhany_nuruu",
	}
}
