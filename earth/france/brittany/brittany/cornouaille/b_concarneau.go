package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔卡诺ConcarneauBarony struct {
	feud.BaseBarony
}

var BConcarneau孔卡诺 feud.Barony = &孔卡诺ConcarneauBarony{}

func init() {
    f := BConcarneau孔卡诺.(*孔卡诺ConcarneauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "concarneau",
		TitleName: "孔卡诺",
		TitleCode: "b_concarneau",
	}
}
