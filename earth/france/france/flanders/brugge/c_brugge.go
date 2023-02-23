package brugge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BruggeCounty interface {
	feud.County
	BAardenburg阿尔登堡() feud.Barony
	BBrugge布鲁日() feud.Barony
	BDamme达默() feud.Barony
	BMale梅尔() feud.Barony
	BNieuwpoort尼乌波特() feud.Barony
	BOostende奥斯坦德() feud.Barony
	BSluys斯勒伊斯() feud.Barony
	BTorhout托尔豪特() feud.Barony
}

type 布鲁日BruggeCounty struct {
	feud.BaseCounty
	阿尔登堡Aardenburg feud.Barony
	布鲁日Brugge      feud.Barony
	达默Damme        feud.Barony
	梅尔Male         feud.Barony
	尼乌波特Nieuwpoort feud.Barony
	奥斯坦德Oostende   feud.Barony
	斯勒伊斯Sluys      feud.Barony
	托尔豪特Torhout    feud.Barony
}

func (c *布鲁日BruggeCounty) BAardenburg阿尔登堡() feud.Barony {
	return c.阿尔登堡Aardenburg
}

func (c *布鲁日BruggeCounty) BBrugge布鲁日() feud.Barony {
	return c.布鲁日Brugge
}

func (c *布鲁日BruggeCounty) BDamme达默() feud.Barony {
	return c.达默Damme
}

func (c *布鲁日BruggeCounty) BMale梅尔() feud.Barony {
	return c.梅尔Male
}

func (c *布鲁日BruggeCounty) BNieuwpoort尼乌波特() feud.Barony {
	return c.尼乌波特Nieuwpoort
}

func (c *布鲁日BruggeCounty) BOostende奥斯坦德() feud.Barony {
	return c.奥斯坦德Oostende
}

func (c *布鲁日BruggeCounty) BSluys斯勒伊斯() feud.Barony {
	return c.斯勒伊斯Sluys
}

func (c *布鲁日BruggeCounty) BTorhout托尔豪特() feud.Barony {
	return c.托尔豪特Torhout
}

var CBrugge布鲁日 BruggeCounty = &布鲁日BruggeCounty{}

func init() {
	f := CBrugge布鲁日.(*布鲁日BruggeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "78",
		Title:     "brugge",
		TitleName: "布鲁日",
		TitleCode: "c_brugge",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔登堡Aardenburg = BAardenburg阿尔登堡
	f.阿尔登堡Aardenburg.SetParent(f)

	f.布鲁日Brugge = BBrugge布鲁日
	f.布鲁日Brugge.SetParent(f)

	f.达默Damme = BDamme达默
	f.达默Damme.SetParent(f)

	f.梅尔Male = BMale梅尔
	f.梅尔Male.SetParent(f)

	f.尼乌波特Nieuwpoort = BNieuwpoort尼乌波特
	f.尼乌波特Nieuwpoort.SetParent(f)

	f.奥斯坦德Oostende = BOostende奥斯坦德
	f.奥斯坦德Oostende.SetParent(f)

	f.斯勒伊斯Sluys = BSluys斯勒伊斯
	f.斯勒伊斯Sluys.SetParent(f)

	f.托尔豪特Torhout = BTorhout托尔豪特
	f.托尔豪特Torhout.SetParent(f)

}
