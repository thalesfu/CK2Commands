package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆霍尔木兹RamhormozBarony struct {
	feud.BaseBarony
}

var BRamhormoz拉姆霍尔木兹 feud.Barony = &拉姆霍尔木兹RamhormozBarony{}

func init() {
	f := BRamhormoz拉姆霍尔木兹.(*拉姆霍尔木兹RamhormozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramhormoz",
		TitleName: "拉姆霍尔木兹",
		TitleCode: "b_ramhormoz",
	}
}
