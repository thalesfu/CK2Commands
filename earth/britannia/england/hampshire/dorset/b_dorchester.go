package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多切斯特DorchesterBarony struct {
	feud.BaseBarony
}

var BDorchester多切斯特 feud.Barony = &多切斯特DorchesterBarony{}

func init() {
	f := BDorchester多切斯特.(*多切斯特DorchesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorchester",
		TitleName: "多切斯特",
		TitleCode: "b_dorchester",
	}
}
