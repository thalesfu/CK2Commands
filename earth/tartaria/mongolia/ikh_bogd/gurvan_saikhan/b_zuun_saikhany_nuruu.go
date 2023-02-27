package gurvan_saikhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 准Zuun_saikhany_nuruuBarony struct {
	feud.BaseBarony
}

var BZuun_saikhany_nuruu准 feud.Barony = &准Zuun_saikhany_nuruuBarony{}

func init() {
    f := BZuun_saikhany_nuruu准.(*准Zuun_saikhany_nuruuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuun_saikhany_nuruu",
		TitleName: "准",
		TitleCode: "b_zuun_saikhany_nuruu",
	}
}
