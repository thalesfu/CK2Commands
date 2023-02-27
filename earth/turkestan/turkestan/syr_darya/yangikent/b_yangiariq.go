package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杨吉尔奇YangiariqBarony struct {
	feud.BaseBarony
}

var BYangiariq杨吉尔奇 feud.Barony = &杨吉尔奇YangiariqBarony{}

func init() {
    f := BYangiariq杨吉尔奇.(*杨吉尔奇YangiariqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yangiariq",
		TitleName: "杨吉尔奇",
		TitleCode: "b_yangiariq",
	}
}
