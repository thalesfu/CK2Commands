package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳卢特NalutBarony struct {
	feud.BaseBarony
}

var BNalut纳卢特 feud.Barony = &纳卢特NalutBarony{}

func init() {
	f := BNalut纳卢特.(*纳卢特NalutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nalut",
		TitleName: "纳卢特",
		TitleCode: "b_nalut",
	}
}
