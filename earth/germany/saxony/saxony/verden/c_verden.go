package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VerdenCounty interface {
    feud.County
    BKirchwalsede基希瓦尔瑟德() 	feud.Barony
    BOttersberg奥特斯贝格() 	feud.Barony
    BOuhusen奥胡森() 	feud.Barony
    BRotenburg_verden罗滕堡() 	feud.Barony
    BSoltau索尔陶() 	feud.Barony
    BVerden费尔登() 	feud.Barony
    BWalsrode瓦尔斯罗德() 	feud.Barony
}

type 费尔登VerdenCounty struct {
	feud.BaseCounty
	基希瓦尔瑟德Kirchwalsede 	feud.Barony
	奥特斯贝格Ottersberg 	feud.Barony
	奥胡森Ouhusen 	feud.Barony
	罗滕堡Rotenburg_verden 	feud.Barony
	索尔陶Soltau 	feud.Barony
	费尔登Verden 	feud.Barony
	瓦尔斯罗德Walsrode 	feud.Barony
}

func (c *费尔登VerdenCounty) BKirchwalsede基希瓦尔瑟德() feud.Barony {
	return c.基希瓦尔瑟德Kirchwalsede
}
    
func (c *费尔登VerdenCounty) BOttersberg奥特斯贝格() feud.Barony {
	return c.奥特斯贝格Ottersberg
}
    
func (c *费尔登VerdenCounty) BOuhusen奥胡森() feud.Barony {
	return c.奥胡森Ouhusen
}
    
func (c *费尔登VerdenCounty) BRotenburg_verden罗滕堡() feud.Barony {
	return c.罗滕堡Rotenburg_verden
}
    
func (c *费尔登VerdenCounty) BSoltau索尔陶() feud.Barony {
	return c.索尔陶Soltau
}
    
func (c *费尔登VerdenCounty) BVerden费尔登() feud.Barony {
	return c.费尔登Verden
}
    
func (c *费尔登VerdenCounty) BWalsrode瓦尔斯罗德() feud.Barony {
	return c.瓦尔斯罗德Walsrode
}
    
var CVerden费尔登 VerdenCounty = &费尔登VerdenCounty{}

func init() {
	f := CVerden费尔登.(*费尔登VerdenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1980",
		Title:     "verden",
		TitleName: "费尔登",
		TitleCode: "c_verden",
		Baronies:  map[string]feud.Barony{},
	}

	f.基希瓦尔瑟德Kirchwalsede = BKirchwalsede基希瓦尔瑟德
	f.基希瓦尔瑟德Kirchwalsede.SetParent(f)
	
	f.奥特斯贝格Ottersberg = BOttersberg奥特斯贝格
	f.奥特斯贝格Ottersberg.SetParent(f)
	
	f.奥胡森Ouhusen = BOuhusen奥胡森
	f.奥胡森Ouhusen.SetParent(f)
	
	f.罗滕堡Rotenburg_verden = BRotenburg_verden罗滕堡
	f.罗滕堡Rotenburg_verden.SetParent(f)
	
	f.索尔陶Soltau = BSoltau索尔陶
	f.索尔陶Soltau.SetParent(f)
	
	f.费尔登Verden = BVerden费尔登
	f.费尔登Verden.SetParent(f)
	
	f.瓦尔斯罗德Walsrode = BWalsrode瓦尔斯罗德
	f.瓦尔斯罗德Walsrode.SetParent(f)
	
}
