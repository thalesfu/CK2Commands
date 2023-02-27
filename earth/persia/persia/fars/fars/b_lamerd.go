package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉默尔德LamerdBarony struct {
	feud.BaseBarony
}

var BLamerd拉默尔德 feud.Barony = &拉默尔德LamerdBarony{}

func init() {
    f := BLamerd拉默尔德.(*拉默尔德LamerdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lamerd",
		TitleName: "拉默尔德",
		TitleCode: "b_lamerd",
	}
}
