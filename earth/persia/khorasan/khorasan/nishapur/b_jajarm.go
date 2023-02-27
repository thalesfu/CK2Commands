package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾贾尔姆JajarmBarony struct {
	feud.BaseBarony
}

var BJajarm贾贾尔姆 feud.Barony = &贾贾尔姆JajarmBarony{}

func init() {
    f := BJajarm贾贾尔姆.(*贾贾尔姆JajarmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jajarm",
		TitleName: "贾贾尔姆",
		TitleCode: "b_jajarm",
	}
}
