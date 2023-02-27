package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeinsterCounty interface {
    feud.County
    BArklow阿克洛() 	feud.Barony
    BCarlow卡洛() 	feud.Barony
    BEnniscorthy恩尼斯科西() 	feud.Barony
    BFerns弗恩斯() 	feud.Barony
    BGlendalough格兰达洛() 	feud.Barony
    BLeighlin利林() 	feud.Barony
    BNaas内斯() 	feud.Barony
    BWexford韦克斯福德() 	feud.Barony
}

type 伦斯特LeinsterCounty struct {
	feud.BaseCounty
	阿克洛Arklow 	feud.Barony
	卡洛Carlow 	feud.Barony
	恩尼斯科西Enniscorthy 	feud.Barony
	弗恩斯Ferns 	feud.Barony
	格兰达洛Glendalough 	feud.Barony
	利林Leighlin 	feud.Barony
	内斯Naas 	feud.Barony
	韦克斯福德Wexford 	feud.Barony
}

func (c *伦斯特LeinsterCounty) BArklow阿克洛() feud.Barony {
	return c.阿克洛Arklow
}
    
func (c *伦斯特LeinsterCounty) BCarlow卡洛() feud.Barony {
	return c.卡洛Carlow
}
    
func (c *伦斯特LeinsterCounty) BEnniscorthy恩尼斯科西() feud.Barony {
	return c.恩尼斯科西Enniscorthy
}
    
func (c *伦斯特LeinsterCounty) BFerns弗恩斯() feud.Barony {
	return c.弗恩斯Ferns
}
    
func (c *伦斯特LeinsterCounty) BGlendalough格兰达洛() feud.Barony {
	return c.格兰达洛Glendalough
}
    
func (c *伦斯特LeinsterCounty) BLeighlin利林() feud.Barony {
	return c.利林Leighlin
}
    
func (c *伦斯特LeinsterCounty) BNaas内斯() feud.Barony {
	return c.内斯Naas
}
    
func (c *伦斯特LeinsterCounty) BWexford韦克斯福德() feud.Barony {
	return c.韦克斯福德Wexford
}
    
var CLeinster伦斯特 LeinsterCounty = &伦斯特LeinsterCounty{}

func init() {
	f := CLeinster伦斯特.(*伦斯特LeinsterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "16",
		Title:     "leinster",
		TitleName: "伦斯特",
		TitleCode: "c_leinster",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克洛Arklow = BArklow阿克洛
	f.阿克洛Arklow.SetParent(f)
	
	f.卡洛Carlow = BCarlow卡洛
	f.卡洛Carlow.SetParent(f)
	
	f.恩尼斯科西Enniscorthy = BEnniscorthy恩尼斯科西
	f.恩尼斯科西Enniscorthy.SetParent(f)
	
	f.弗恩斯Ferns = BFerns弗恩斯
	f.弗恩斯Ferns.SetParent(f)
	
	f.格兰达洛Glendalough = BGlendalough格兰达洛
	f.格兰达洛Glendalough.SetParent(f)
	
	f.利林Leighlin = BLeighlin利林
	f.利林Leighlin.SetParent(f)
	
	f.内斯Naas = BNaas内斯
	f.内斯Naas.SetParent(f)
	
	f.韦克斯福德Wexford = BWexford韦克斯福德
	f.韦克斯福德Wexford.SetParent(f)
	
}
