package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣蓬德托米耶尔StponsdethomieresBarony struct {
	feud.BaseBarony
}

var BStponsdethomieres圣蓬德托米耶尔 feud.Barony = &圣蓬德托米耶尔StponsdethomieresBarony{}

func init() {
    f := BStponsdethomieres圣蓬德托米耶尔.(*圣蓬德托米耶尔StponsdethomieresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stponsdethomieres",
		TitleName: "圣蓬德托米耶尔",
		TitleCode: "b_stponsdethomieres",
	}
}
