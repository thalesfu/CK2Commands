package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊德阿米Id_amyBarony struct {
	feud.BaseBarony
}

var BId_amy伊德阿米 feud.Barony = &伊德阿米Id_amyBarony{}

func init() {
    f := BId_amy伊德阿米.(*伊德阿米Id_amyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "id_amy",
		TitleName: "伊德阿米",
		TitleCode: "b_id_amy",
	}
}
