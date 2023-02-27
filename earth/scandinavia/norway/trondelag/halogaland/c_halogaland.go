package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HalogalandCounty interface {
    feud.County
    BAlstahaug阿尔斯塔海于格() 	feud.Barony
    BBindal宾达() 	feud.Barony
    BBrunnoy布吕内于() 	feud.Barony
    BHattfjelldalen哈特杰达冷() 	feud.Barony
    BLein勒因() 	feud.Barony
    BMosjoen莫舍恩() 	feud.Barony
    BSomna瑟姆纳() 	feud.Barony
    BVeiga韦加() 	feud.Barony
}

type 霍洛加兰HalogalandCounty struct {
	feud.BaseCounty
	阿尔斯塔海于格Alstahaug 	feud.Barony
	宾达Bindal 	feud.Barony
	布吕内于Brunnoy 	feud.Barony
	哈特杰达冷Hattfjelldalen 	feud.Barony
	勒因Lein 	feud.Barony
	莫舍恩Mosjoen 	feud.Barony
	瑟姆纳Somna 	feud.Barony
	韦加Veiga 	feud.Barony
}

func (c *霍洛加兰HalogalandCounty) BAlstahaug阿尔斯塔海于格() feud.Barony {
	return c.阿尔斯塔海于格Alstahaug
}
    
func (c *霍洛加兰HalogalandCounty) BBindal宾达() feud.Barony {
	return c.宾达Bindal
}
    
func (c *霍洛加兰HalogalandCounty) BBrunnoy布吕内于() feud.Barony {
	return c.布吕内于Brunnoy
}
    
func (c *霍洛加兰HalogalandCounty) BHattfjelldalen哈特杰达冷() feud.Barony {
	return c.哈特杰达冷Hattfjelldalen
}
    
func (c *霍洛加兰HalogalandCounty) BLein勒因() feud.Barony {
	return c.勒因Lein
}
    
func (c *霍洛加兰HalogalandCounty) BMosjoen莫舍恩() feud.Barony {
	return c.莫舍恩Mosjoen
}
    
func (c *霍洛加兰HalogalandCounty) BSomna瑟姆纳() feud.Barony {
	return c.瑟姆纳Somna
}
    
func (c *霍洛加兰HalogalandCounty) BVeiga韦加() feud.Barony {
	return c.韦加Veiga
}
    
var CHalogaland霍洛加兰 HalogalandCounty = &霍洛加兰HalogalandCounty{}

func init() {
	f := CHalogaland霍洛加兰.(*霍洛加兰HalogalandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "278",
		Title:     "halogaland",
		TitleName: "霍洛加兰",
		TitleCode: "c_halogaland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯塔海于格Alstahaug = BAlstahaug阿尔斯塔海于格
	f.阿尔斯塔海于格Alstahaug.SetParent(f)
	
	f.宾达Bindal = BBindal宾达
	f.宾达Bindal.SetParent(f)
	
	f.布吕内于Brunnoy = BBrunnoy布吕内于
	f.布吕内于Brunnoy.SetParent(f)
	
	f.哈特杰达冷Hattfjelldalen = BHattfjelldalen哈特杰达冷
	f.哈特杰达冷Hattfjelldalen.SetParent(f)
	
	f.勒因Lein = BLein勒因
	f.勒因Lein.SetParent(f)
	
	f.莫舍恩Mosjoen = BMosjoen莫舍恩
	f.莫舍恩Mosjoen.SetParent(f)
	
	f.瑟姆纳Somna = BSomna瑟姆纳
	f.瑟姆纳Somna.SetParent(f)
	
	f.韦加Veiga = BVeiga韦加
	f.韦加Veiga.SetParent(f)
	
}
