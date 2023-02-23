package brittany

import (
	"github.com/thalesfu/CK2Commands/earth/france/brittany/brittany/cornouaille"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/brittany/poher"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/brittany/vannes"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrittanyDuke interface {
	feud.Duke
	CCornouaille科努瓦耶() cornouaille.CornouailleCounty
	CPoher波埃() poher.PoherCounty
	CVannes瓦讷() vannes.VannesCounty
}

type 布列塔尼BrittanyDuke struct {
	feud.BaseDuke
	科努瓦耶Cornouaille cornouaille.CornouailleCounty
	波埃Poher         poher.PoherCounty
	瓦讷Vannes        vannes.VannesCounty
}

func (d *布列塔尼BrittanyDuke) CCornouaille科努瓦耶() cornouaille.CornouailleCounty {
	return d.科努瓦耶Cornouaille
}

func (d *布列塔尼BrittanyDuke) CPoher波埃() poher.PoherCounty {
	return d.波埃Poher
}

func (d *布列塔尼BrittanyDuke) CVannes瓦讷() vannes.VannesCounty {
	return d.瓦讷Vannes
}

var DBrittany布列塔尼 BrittanyDuke = &布列塔尼BrittanyDuke{}

func init() {
	f := DBrittany布列塔尼.(*布列塔尼BrittanyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "brittany",
		TitleName: "布列塔尼",
		TitleCode: "d_brittany",
		Counties:  map[string]feud.County{},
	}

	f.科努瓦耶Cornouaille = cornouaille.CCornouaille科努瓦耶
	f.科努瓦耶Cornouaille.SetParent(f)

	f.波埃Poher = poher.CPoher波埃
	f.波埃Poher.SetParent(f)

	f.瓦讷Vannes = vannes.CVannes瓦讷
	f.瓦讷Vannes.SetParent(f)

}
