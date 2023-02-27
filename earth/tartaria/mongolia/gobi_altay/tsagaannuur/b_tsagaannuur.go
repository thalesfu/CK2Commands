package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查干诺尔TsagaannuurBarony struct {
	feud.BaseBarony
}

var BTsagaannuur查干诺尔 feud.Barony = &查干诺尔TsagaannuurBarony{}

func init() {
    f := BTsagaannuur查干诺尔.(*查干诺尔TsagaannuurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsagaannuur",
		TitleName: "查干诺尔",
		TitleCode: "b_tsagaannuur",
	}
}
