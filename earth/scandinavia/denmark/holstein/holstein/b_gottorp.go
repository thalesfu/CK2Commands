package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈托普GottorpBarony struct {
	feud.BaseBarony
}

var BGottorp戈托普 feud.Barony = &戈托普GottorpBarony{}

func init() {
    f := BGottorp戈托普.(*戈托普GottorpBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gottorp",
		TitleName: "戈托普",
		TitleCode: "b_gottorp",
	}
}
