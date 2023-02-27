package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀曼那护罗DamanhurBarony struct {
	feud.BaseBarony
}

var BDamanhur陀曼那护罗 feud.Barony = &陀曼那护罗DamanhurBarony{}

func init() {
    f := BDamanhur陀曼那护罗.(*陀曼那护罗DamanhurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damanhur",
		TitleName: "陀曼那护罗",
		TitleCode: "b_damanhur",
	}
}
