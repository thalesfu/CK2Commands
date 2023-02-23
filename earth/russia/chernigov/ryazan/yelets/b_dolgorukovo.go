package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尔戈鲁科沃DolgorukovoBarony struct {
	feud.BaseBarony
}

var BDolgorukovo多尔戈鲁科沃 feud.Barony = &多尔戈鲁科沃DolgorukovoBarony{}

func init() {
	f := BDolgorukovo多尔戈鲁科沃.(*多尔戈鲁科沃DolgorukovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolgorukovo",
		TitleName: "多尔戈鲁科沃",
		TitleCode: "b_dolgorukovo",
	}
}
