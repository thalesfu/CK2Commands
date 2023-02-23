package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桐冲岔TongchongchaBarony struct {
	feud.BaseBarony
}

var BTongchongcha桐冲岔 feud.Barony = &桐冲岔TongchongchaBarony{}

func init() {
	f := BTongchongcha桐冲岔.(*桐冲岔TongchongchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tongchongcha",
		TitleName: "桐冲岔",
		TitleCode: "b_tongchongcha",
	}
}
