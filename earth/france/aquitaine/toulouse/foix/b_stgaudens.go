package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣戈当StgaudensBarony struct {
	feud.BaseBarony
}

var BStgaudens圣戈当 feud.Barony = &圣戈当StgaudensBarony{}

func init() {
    f := BStgaudens圣戈当.(*圣戈当StgaudensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stgaudens",
		TitleName: "圣戈当",
		TitleCode: "b_stgaudens",
	}
}
