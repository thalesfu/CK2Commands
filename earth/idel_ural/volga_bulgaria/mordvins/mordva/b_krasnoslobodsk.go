package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯诺斯洛博茨克KrasnoslobodskBarony struct {
	feud.BaseBarony
}

var BKrasnoslobodsk克拉斯诺斯洛博茨克 feud.Barony = &克拉斯诺斯洛博茨克KrasnoslobodskBarony{}

func init() {
    f := BKrasnoslobodsk克拉斯诺斯洛博茨克.(*克拉斯诺斯洛博茨克KrasnoslobodskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnoslobodsk",
		TitleName: "克拉斯诺斯洛博茨克",
		TitleCode: "b_krasnoslobodsk",
	}
}
