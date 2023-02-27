package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰特德尔马埃斯特雷FuentedelmaestreBarony struct {
	feud.BaseBarony
}

var BFuentedelmaestre丰特德尔马埃斯特雷 feud.Barony = &丰特德尔马埃斯特雷FuentedelmaestreBarony{}

func init() {
    f := BFuentedelmaestre丰特德尔马埃斯特雷.(*丰特德尔马埃斯特雷FuentedelmaestreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuentedelmaestre",
		TitleName: "丰特德尔马埃斯特雷",
		TitleCode: "b_fuentedelmaestre",
	}
}
