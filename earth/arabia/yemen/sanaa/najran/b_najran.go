package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳季兰NajranBarony struct {
	feud.BaseBarony
}

var BNajran纳季兰 feud.Barony = &纳季兰NajranBarony{}

func init() {
	f := BNajran纳季兰.(*纳季兰NajranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "najran",
		TitleName: "纳季兰",
		TitleCode: "b_najran",
	}
}
