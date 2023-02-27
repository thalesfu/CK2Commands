package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔拉哈GuerrahBarony struct {
	feud.BaseBarony
}

var BGuerrah古尔拉哈 feud.Barony = &古尔拉哈GuerrahBarony{}

func init() {
    f := BGuerrah古尔拉哈.(*古尔拉哈GuerrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guerrah",
		TitleName: "古尔拉哈",
		TitleCode: "b_guerrah",
	}
}
