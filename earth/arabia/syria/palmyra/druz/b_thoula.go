package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特霍拉ThoulaBarony struct {
	feud.BaseBarony
}

var BThoula特霍拉 feud.Barony = &特霍拉ThoulaBarony{}

func init() {
    f := BThoula特霍拉.(*特霍拉ThoulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thoula",
		TitleName: "特霍拉",
		TitleCode: "b_thoula",
	}
}
