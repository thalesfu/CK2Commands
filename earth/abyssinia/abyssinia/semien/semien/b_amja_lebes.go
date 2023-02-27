package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆贾勒贝斯Amja_lebesBarony struct {
	feud.BaseBarony
}

var BAmja_lebes阿姆贾勒贝斯 feud.Barony = &阿姆贾勒贝斯Amja_lebesBarony{}

func init() {
    f := BAmja_lebes阿姆贾勒贝斯.(*阿姆贾勒贝斯Amja_lebesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amja_lebes",
		TitleName: "阿姆贾勒贝斯",
		TitleCode: "b_amja_lebes",
	}
}
