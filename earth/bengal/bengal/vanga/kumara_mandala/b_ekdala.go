package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夷迦陀罗EkdalaBarony struct {
	feud.BaseBarony
}

var BEkdala夷迦陀罗 feud.Barony = &夷迦陀罗EkdalaBarony{}

func init() {
    f := BEkdala夷迦陀罗.(*夷迦陀罗EkdalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ekdala",
		TitleName: "夷迦陀罗",
		TitleCode: "b_ekdala",
	}
}
