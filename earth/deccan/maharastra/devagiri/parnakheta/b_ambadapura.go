package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菴婆陀补罗AmbadapuraBarony struct {
	feud.BaseBarony
}

var BAmbadapura菴婆陀补罗 feud.Barony = &菴婆陀补罗AmbadapuraBarony{}

func init() {
    f := BAmbadapura菴婆陀补罗.(*菴婆陀补罗AmbadapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ambadapura",
		TitleName: "菴婆陀补罗",
		TitleCode: "b_ambadapura",
	}
}
