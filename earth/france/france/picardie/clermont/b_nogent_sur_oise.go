package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺让Nogent_sur_oiseBarony struct {
	feud.BaseBarony
}

var BNogent_sur_oise诺让 feud.Barony = &诺让Nogent_sur_oiseBarony{}

func init() {
    f := BNogent_sur_oise诺让.(*诺让Nogent_sur_oiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogent_sur_oise",
		TitleName: "诺让",
		TitleCode: "b_nogent_sur_oise",
	}
}
