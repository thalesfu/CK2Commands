package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜廷赫姆DuetinghemBarony struct {
	feud.BaseBarony
}

var BDuetinghem杜廷赫姆 feud.Barony = &杜廷赫姆DuetinghemBarony{}

func init() {
    f := BDuetinghem杜廷赫姆.(*杜廷赫姆DuetinghemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duetinghem",
		TitleName: "杜廷赫姆",
		TitleCode: "b_duetinghem",
	}
}
