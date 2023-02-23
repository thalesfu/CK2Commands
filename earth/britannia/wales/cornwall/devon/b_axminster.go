package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克斯明斯特AxminsterBarony struct {
	feud.BaseBarony
}

var BAxminster阿克斯明斯特 feud.Barony = &阿克斯明斯特AxminsterBarony{}

func init() {
	f := BAxminster阿克斯明斯特.(*阿克斯明斯特AxminsterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "axminster",
		TitleName: "阿克斯明斯特",
		TitleCode: "b_axminster",
	}
}
