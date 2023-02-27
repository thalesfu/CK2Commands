package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因贝尼迈特海尔AinbenimatharBarony struct {
	feud.BaseBarony
}

var BAinbenimathar艾因贝尼迈特海尔 feud.Barony = &艾因贝尼迈特海尔AinbenimatharBarony{}

func init() {
    f := BAinbenimathar艾因贝尼迈特海尔.(*艾因贝尼迈特海尔AinbenimatharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ainbenimathar",
		TitleName: "艾因贝尼迈特海尔",
		TitleCode: "b_ainbenimathar",
	}
}
