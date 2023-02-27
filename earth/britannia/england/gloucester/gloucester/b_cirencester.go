package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛伦塞斯特CirencesterBarony struct {
	feud.BaseBarony
}

var BCirencester赛伦塞斯特 feud.Barony = &赛伦塞斯特CirencesterBarony{}

func init() {
    f := BCirencester赛伦塞斯特.(*赛伦塞斯特CirencesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cirencester",
		TitleName: "赛伦塞斯特",
		TitleCode: "b_cirencester",
	}
}
