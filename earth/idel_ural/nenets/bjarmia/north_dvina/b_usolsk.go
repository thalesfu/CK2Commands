package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌索利斯克UsolskBarony struct {
	feud.BaseBarony
}

var BUsolsk乌索利斯克 feud.Barony = &乌索利斯克UsolskBarony{}

func init() {
    f := BUsolsk乌索利斯克.(*乌索利斯克UsolskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usolsk",
		TitleName: "乌索利斯克",
		TitleCode: "b_usolsk",
	}
}
