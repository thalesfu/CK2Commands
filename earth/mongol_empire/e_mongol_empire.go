package mongol_empire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Mongol_empireEmpire interface {
    feud.Empire
}

type 大蒙古国Mongol_empireEmpire struct {
	feud.BaseEmpire
}

var EMongol_empire大蒙古国 Mongol_empireEmpire = &大蒙古国Mongol_empireEmpire{}

func init() {
	f := EMongol_empire大蒙古国.(*大蒙古国Mongol_empireEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "mongol_empire",
		TitleName: "大蒙古国",
		TitleCode: "e_mongol_empire",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
