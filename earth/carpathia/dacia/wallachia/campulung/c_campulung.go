package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CampulungCounty interface {
    feud.County
    BArges阿尔杰什() 	feud.Barony
    BBovora博罗瓦() 	feud.Barony
    BBran布兰() 	feud.Barony
    BCampulung肯普隆格() 	feud.Barony
    BCozla科兹拉() 	feud.Barony
    BMioveni苗韦尼() 	feud.Barony
    BPitesti皮特什蒂() 	feud.Barony
}

type 肯普隆格CampulungCounty struct {
	feud.BaseCounty
	阿尔杰什Arges 	feud.Barony
	博罗瓦Bovora 	feud.Barony
	布兰Bran 	feud.Barony
	肯普隆格Campulung 	feud.Barony
	科兹拉Cozla 	feud.Barony
	苗韦尼Mioveni 	feud.Barony
	皮特什蒂Pitesti 	feud.Barony
}

func (c *肯普隆格CampulungCounty) BArges阿尔杰什() feud.Barony {
	return c.阿尔杰什Arges
}
    
func (c *肯普隆格CampulungCounty) BBovora博罗瓦() feud.Barony {
	return c.博罗瓦Bovora
}
    
func (c *肯普隆格CampulungCounty) BBran布兰() feud.Barony {
	return c.布兰Bran
}
    
func (c *肯普隆格CampulungCounty) BCampulung肯普隆格() feud.Barony {
	return c.肯普隆格Campulung
}
    
func (c *肯普隆格CampulungCounty) BCozla科兹拉() feud.Barony {
	return c.科兹拉Cozla
}
    
func (c *肯普隆格CampulungCounty) BMioveni苗韦尼() feud.Barony {
	return c.苗韦尼Mioveni
}
    
func (c *肯普隆格CampulungCounty) BPitesti皮特什蒂() feud.Barony {
	return c.皮特什蒂Pitesti
}
    
var CCampulung肯普隆格 CampulungCounty = &肯普隆格CampulungCounty{}

func init() {
	f := CCampulung肯普隆格.(*肯普隆格CampulungCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1644",
		Title:     "campulung",
		TitleName: "肯普隆格",
		TitleCode: "c_campulung",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔杰什Arges = BArges阿尔杰什
	f.阿尔杰什Arges.SetParent(f)
	
	f.博罗瓦Bovora = BBovora博罗瓦
	f.博罗瓦Bovora.SetParent(f)
	
	f.布兰Bran = BBran布兰
	f.布兰Bran.SetParent(f)
	
	f.肯普隆格Campulung = BCampulung肯普隆格
	f.肯普隆格Campulung.SetParent(f)
	
	f.科兹拉Cozla = BCozla科兹拉
	f.科兹拉Cozla.SetParent(f)
	
	f.苗韦尼Mioveni = BMioveni苗韦尼
	f.苗韦尼Mioveni.SetParent(f)
	
	f.皮特什蒂Pitesti = BPitesti皮特什蒂
	f.皮特什蒂Pitesti.SetParent(f)
	
}
