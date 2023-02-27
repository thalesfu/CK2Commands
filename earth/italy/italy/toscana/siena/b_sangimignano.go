package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣吉米尼亚诺SangimignanoBarony struct {
	feud.BaseBarony
}

var BSangimignano圣吉米尼亚诺 feud.Barony = &圣吉米尼亚诺SangimignanoBarony{}

func init() {
    f := BSangimignano圣吉米尼亚诺.(*圣吉米尼亚诺SangimignanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangimignano",
		TitleName: "圣吉米尼亚诺",
		TitleCode: "b_sangimignano",
	}
}
