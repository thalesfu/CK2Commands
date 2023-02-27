package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿杰富尔Al_ajfurBarony struct {
	feud.BaseBarony
}

var BAl_ajfur阿杰富尔 feud.Barony = &阿杰富尔Al_ajfurBarony{}

func init() {
    f := BAl_ajfur阿杰富尔.(*阿杰富尔Al_ajfurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_ajfur",
		TitleName: "阿杰富尔",
		TitleCode: "b_al_ajfur",
	}
}
