package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾格巴巴JagbadaBarony struct {
	feud.BaseBarony
}

var BJagbada贾格巴巴 feud.Barony = &贾格巴巴JagbadaBarony{}

func init() {
	f := BJagbada贾格巴巴.(*贾格巴巴JagbadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jagbada",
		TitleName: "贾格巴巴",
		TitleCode: "b_jagbada",
	}
}
