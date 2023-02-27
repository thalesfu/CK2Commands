package bizerte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BizerteCounty interface {
    feud.County
    BAlmatin艾玛汀() 	feud.Barony
    BBizerte比塞大() 	feud.Barony
    BGhezala格赫扎拉() 	feud.Barony
    BJendouba坚杜拜() 	feud.Barony
    BManziljamal曼济扎马拉() 	feud.Barony
    BNetza奈茨阿() 	feud.Barony
    BTamra塔玛拉() 	feud.Barony
    BTunbeja巴杰() 	feud.Barony
}

type 比塞大BizerteCounty struct {
	feud.BaseCounty
	艾玛汀Almatin 	feud.Barony
	比塞大Bizerte 	feud.Barony
	格赫扎拉Ghezala 	feud.Barony
	坚杜拜Jendouba 	feud.Barony
	曼济扎马拉Manziljamal 	feud.Barony
	奈茨阿Netza 	feud.Barony
	塔玛拉Tamra 	feud.Barony
	巴杰Tunbeja 	feud.Barony
}

func (c *比塞大BizerteCounty) BAlmatin艾玛汀() feud.Barony {
	return c.艾玛汀Almatin
}
    
func (c *比塞大BizerteCounty) BBizerte比塞大() feud.Barony {
	return c.比塞大Bizerte
}
    
func (c *比塞大BizerteCounty) BGhezala格赫扎拉() feud.Barony {
	return c.格赫扎拉Ghezala
}
    
func (c *比塞大BizerteCounty) BJendouba坚杜拜() feud.Barony {
	return c.坚杜拜Jendouba
}
    
func (c *比塞大BizerteCounty) BManziljamal曼济扎马拉() feud.Barony {
	return c.曼济扎马拉Manziljamal
}
    
func (c *比塞大BizerteCounty) BNetza奈茨阿() feud.Barony {
	return c.奈茨阿Netza
}
    
func (c *比塞大BizerteCounty) BTamra塔玛拉() feud.Barony {
	return c.塔玛拉Tamra
}
    
func (c *比塞大BizerteCounty) BTunbeja巴杰() feud.Barony {
	return c.巴杰Tunbeja
}
    
var CBizerte比塞大 BizerteCounty = &比塞大BizerteCounty{}

func init() {
	f := CBizerte比塞大.(*比塞大BizerteCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "819",
		Title:     "bizerte",
		TitleName: "比塞大",
		TitleCode: "c_bizerte",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾玛汀Almatin = BAlmatin艾玛汀
	f.艾玛汀Almatin.SetParent(f)
	
	f.比塞大Bizerte = BBizerte比塞大
	f.比塞大Bizerte.SetParent(f)
	
	f.格赫扎拉Ghezala = BGhezala格赫扎拉
	f.格赫扎拉Ghezala.SetParent(f)
	
	f.坚杜拜Jendouba = BJendouba坚杜拜
	f.坚杜拜Jendouba.SetParent(f)
	
	f.曼济扎马拉Manziljamal = BManziljamal曼济扎马拉
	f.曼济扎马拉Manziljamal.SetParent(f)
	
	f.奈茨阿Netza = BNetza奈茨阿
	f.奈茨阿Netza.SetParent(f)
	
	f.塔玛拉Tamra = BTamra塔玛拉
	f.塔玛拉Tamra.SetParent(f)
	
	f.巴杰Tunbeja = BTunbeja巴杰
	f.巴杰Tunbeja.SetParent(f)
	
}
