package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳恩迪亚尔NandyalBarony struct {
	feud.BaseBarony
}

var BNandyal纳恩迪亚尔 feud.Barony = &纳恩迪亚尔NandyalBarony{}

func init() {
	f := BNandyal纳恩迪亚尔.(*纳恩迪亚尔NandyalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandyal",
		TitleName: "纳恩迪亚尔",
		TitleCode: "b_nandyal",
	}
}
