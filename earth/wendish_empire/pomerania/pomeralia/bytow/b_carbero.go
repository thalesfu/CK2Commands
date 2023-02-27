package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔贝罗CarberoBarony struct {
	feud.BaseBarony
}

var BCarbero卡尔贝罗 feud.Barony = &卡尔贝罗CarberoBarony{}

func init() {
    f := BCarbero卡尔贝罗.(*卡尔贝罗CarberoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carbero",
		TitleName: "卡尔贝罗",
		TitleCode: "b_carbero",
	}
}
