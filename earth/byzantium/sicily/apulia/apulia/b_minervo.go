package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米内尔沃MinervoBarony struct {
	feud.BaseBarony
}

var BMinervo米内尔沃 feud.Barony = &米内尔沃MinervoBarony{}

func init() {
    f := BMinervo米内尔沃.(*米内尔沃MinervoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minervo",
		TitleName: "米内尔沃",
		TitleCode: "b_minervo",
	}
}
