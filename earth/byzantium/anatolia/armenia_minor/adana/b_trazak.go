package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉扎克TrazakBarony struct {
	feud.BaseBarony
}

var BTrazak特拉扎克 feud.Barony = &特拉扎克TrazakBarony{}

func init() {
    f := BTrazak特拉扎克.(*特拉扎克TrazakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trazak",
		TitleName: "特拉扎克",
		TitleCode: "b_trazak",
	}
}
