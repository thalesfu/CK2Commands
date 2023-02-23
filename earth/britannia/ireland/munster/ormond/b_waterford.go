package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃特福德WaterfordBarony struct {
	feud.BaseBarony
}

var BWaterford沃特福德 feud.Barony = &沃特福德WaterfordBarony{}

func init() {
	f := BWaterford沃特福德.(*沃特福德WaterfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waterford",
		TitleName: "沃特福德",
		TitleCode: "b_waterford",
	}
}
