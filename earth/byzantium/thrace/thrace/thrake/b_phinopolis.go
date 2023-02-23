package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲诺波利斯PhinopolisBarony struct {
	feud.BaseBarony
}

var BPhinopolis菲诺波利斯 feud.Barony = &菲诺波利斯PhinopolisBarony{}

func init() {
	f := BPhinopolis菲诺波利斯.(*菲诺波利斯PhinopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phinopolis",
		TitleName: "菲诺波利斯",
		TitleCode: "b_phinopolis",
	}
}
