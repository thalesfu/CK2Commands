package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃特福德WatfordBarony struct {
	feud.BaseBarony
}

var BWatford沃特福德 feud.Barony = &沃特福德WatfordBarony{}

func init() {
    f := BWatford沃特福德.(*沃特福德WatfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "watford",
		TitleName: "沃特福德",
		TitleCode: "b_watford",
	}
}
