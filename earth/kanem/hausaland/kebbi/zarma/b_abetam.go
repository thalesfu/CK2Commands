package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿贝塔姆AbetamBarony struct {
	feud.BaseBarony
}

var BAbetam阿贝塔姆 feud.Barony = &阿贝塔姆AbetamBarony{}

func init() {
    f := BAbetam阿贝塔姆.(*阿贝塔姆AbetamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abetam",
		TitleName: "阿贝塔姆",
		TitleCode: "b_abetam",
	}
}
