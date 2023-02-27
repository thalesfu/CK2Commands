package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列那Lena_bolshayaBarony struct {
	feud.BaseBarony
}

var BLena_bolshaya列那 feud.Barony = &列那Lena_bolshayaBarony{}

func init() {
    f := BLena_bolshaya列那.(*列那Lena_bolshayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lena_bolshaya",
		TitleName: "列那",
		TitleCode: "b_lena_bolshaya",
	}
}
