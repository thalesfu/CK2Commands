package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施拉沃SchlaweBarony struct {
	feud.BaseBarony
}

var BSchlawe施拉沃 feud.Barony = &施拉沃SchlaweBarony{}

func init() {
    f := BSchlawe施拉沃.(*施拉沃SchlaweBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schlawe",
		TitleName: "施拉沃",
		TitleCode: "b_schlawe",
	}
}
