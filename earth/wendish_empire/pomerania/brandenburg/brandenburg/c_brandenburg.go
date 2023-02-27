package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrandenburgCounty interface {
    feud.County
    BBelzig贝尔齐希() 	feud.Barony
    BBerlin柏林() 	feud.Barony
    BBrandenburg勃兰登堡() 	feud.Barony
    BJuterborg于特博格() 	feud.Barony
    BMuncheberg明谢贝格() 	feud.Barony
    BOranienburg奥拉宁堡() 	feud.Barony
    BRuppin鲁平() 	feud.Barony
}

type 勃兰登堡BrandenburgCounty struct {
	feud.BaseCounty
	贝尔齐希Belzig 	feud.Barony
	柏林Berlin 	feud.Barony
	勃兰登堡Brandenburg 	feud.Barony
	于特博格Juterborg 	feud.Barony
	明谢贝格Muncheberg 	feud.Barony
	奥拉宁堡Oranienburg 	feud.Barony
	鲁平Ruppin 	feud.Barony
}

func (c *勃兰登堡BrandenburgCounty) BBelzig贝尔齐希() feud.Barony {
	return c.贝尔齐希Belzig
}
    
func (c *勃兰登堡BrandenburgCounty) BBerlin柏林() feud.Barony {
	return c.柏林Berlin
}
    
func (c *勃兰登堡BrandenburgCounty) BBrandenburg勃兰登堡() feud.Barony {
	return c.勃兰登堡Brandenburg
}
    
func (c *勃兰登堡BrandenburgCounty) BJuterborg于特博格() feud.Barony {
	return c.于特博格Juterborg
}
    
func (c *勃兰登堡BrandenburgCounty) BMuncheberg明谢贝格() feud.Barony {
	return c.明谢贝格Muncheberg
}
    
func (c *勃兰登堡BrandenburgCounty) BOranienburg奥拉宁堡() feud.Barony {
	return c.奥拉宁堡Oranienburg
}
    
func (c *勃兰登堡BrandenburgCounty) BRuppin鲁平() feud.Barony {
	return c.鲁平Ruppin
}
    
var CBrandenburg勃兰登堡 BrandenburgCounty = &勃兰登堡BrandenburgCounty{}

func init() {
	f := CBrandenburg勃兰登堡.(*勃兰登堡BrandenburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "365",
		Title:     "brandenburg",
		TitleName: "勃兰登堡",
		TitleCode: "c_brandenburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔齐希Belzig = BBelzig贝尔齐希
	f.贝尔齐希Belzig.SetParent(f)
	
	f.柏林Berlin = BBerlin柏林
	f.柏林Berlin.SetParent(f)
	
	f.勃兰登堡Brandenburg = BBrandenburg勃兰登堡
	f.勃兰登堡Brandenburg.SetParent(f)
	
	f.于特博格Juterborg = BJuterborg于特博格
	f.于特博格Juterborg.SetParent(f)
	
	f.明谢贝格Muncheberg = BMuncheberg明谢贝格
	f.明谢贝格Muncheberg.SetParent(f)
	
	f.奥拉宁堡Oranienburg = BOranienburg奥拉宁堡
	f.奥拉宁堡Oranienburg.SetParent(f)
	
	f.鲁平Ruppin = BRuppin鲁平
	f.鲁平Ruppin.SetParent(f)
	
}
