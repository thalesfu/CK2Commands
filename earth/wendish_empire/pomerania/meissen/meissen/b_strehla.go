package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施特雷拉StrehlaBarony struct {
	feud.BaseBarony
}

var BStrehla施特雷拉 feud.Barony = &施特雷拉StrehlaBarony{}

func init() {
    f := BStrehla施特雷拉.(*施特雷拉StrehlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strehla",
		TitleName: "施特雷拉",
		TitleCode: "b_strehla",
	}
}
