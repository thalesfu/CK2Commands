package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙泰居MontaiguBarony struct {
	feud.BaseBarony
}

var BMontaigu蒙泰居 feud.Barony = &蒙泰居MontaiguBarony{}

func init() {
    f := BMontaigu蒙泰居.(*蒙泰居MontaiguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montaigu",
		TitleName: "蒙泰居",
		TitleCode: "b_montaigu",
	}
}
