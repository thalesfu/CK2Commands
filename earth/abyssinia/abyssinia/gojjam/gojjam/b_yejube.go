package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶朱贝YejubeBarony struct {
	feud.BaseBarony
}

var BYejube耶朱贝 feud.Barony = &耶朱贝YejubeBarony{}

func init() {
    f := BYejube耶朱贝.(*耶朱贝YejubeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yejube",
		TitleName: "耶朱贝",
		TitleCode: "b_yejube",
	}
}
