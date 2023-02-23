package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈姆拉AlhamraBarony struct {
	feud.BaseBarony
}

var BAlhamra哈姆拉 feud.Barony = &哈姆拉AlhamraBarony{}

func init() {
	f := BAlhamra哈姆拉.(*哈姆拉AlhamraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhamra",
		TitleName: "哈姆拉",
		TitleCode: "b_alhamra",
	}
}
