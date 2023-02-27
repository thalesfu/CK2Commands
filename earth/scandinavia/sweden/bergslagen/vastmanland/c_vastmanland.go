package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VastmanlandCounty interface {
    feud.County
    BArboga阿尔博加() 	feud.Barony
    BBadelunda巴德伦达() 	feud.Barony
    BKoping雪平() 	feud.Barony
    BKopingshus雪平斯胡斯() 	feud.Barony
    BMunktorp蒙克托普() 	feud.Barony
    BNorberg努尔贝里() 	feud.Barony
    BSkinnskatteberg欣斯卡特贝里() 	feud.Barony
    BVasteras韦斯特罗斯() 	feud.Barony
}

type 西曼兰VastmanlandCounty struct {
	feud.BaseCounty
	阿尔博加Arboga 	feud.Barony
	巴德伦达Badelunda 	feud.Barony
	雪平Koping 	feud.Barony
	雪平斯胡斯Kopingshus 	feud.Barony
	蒙克托普Munktorp 	feud.Barony
	努尔贝里Norberg 	feud.Barony
	欣斯卡特贝里Skinnskatteberg 	feud.Barony
	韦斯特罗斯Vasteras 	feud.Barony
}

func (c *西曼兰VastmanlandCounty) BArboga阿尔博加() feud.Barony {
	return c.阿尔博加Arboga
}
    
func (c *西曼兰VastmanlandCounty) BBadelunda巴德伦达() feud.Barony {
	return c.巴德伦达Badelunda
}
    
func (c *西曼兰VastmanlandCounty) BKoping雪平() feud.Barony {
	return c.雪平Koping
}
    
func (c *西曼兰VastmanlandCounty) BKopingshus雪平斯胡斯() feud.Barony {
	return c.雪平斯胡斯Kopingshus
}
    
func (c *西曼兰VastmanlandCounty) BMunktorp蒙克托普() feud.Barony {
	return c.蒙克托普Munktorp
}
    
func (c *西曼兰VastmanlandCounty) BNorberg努尔贝里() feud.Barony {
	return c.努尔贝里Norberg
}
    
func (c *西曼兰VastmanlandCounty) BSkinnskatteberg欣斯卡特贝里() feud.Barony {
	return c.欣斯卡特贝里Skinnskatteberg
}
    
func (c *西曼兰VastmanlandCounty) BVasteras韦斯特罗斯() feud.Barony {
	return c.韦斯特罗斯Vasteras
}
    
var CVastmanland西曼兰 VastmanlandCounty = &西曼兰VastmanlandCounty{}

func init() {
	f := CVastmanland西曼兰.(*西曼兰VastmanlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "289",
		Title:     "vastmanland",
		TitleName: "西曼兰",
		TitleCode: "c_vastmanland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔博加Arboga = BArboga阿尔博加
	f.阿尔博加Arboga.SetParent(f)
	
	f.巴德伦达Badelunda = BBadelunda巴德伦达
	f.巴德伦达Badelunda.SetParent(f)
	
	f.雪平Koping = BKoping雪平
	f.雪平Koping.SetParent(f)
	
	f.雪平斯胡斯Kopingshus = BKopingshus雪平斯胡斯
	f.雪平斯胡斯Kopingshus.SetParent(f)
	
	f.蒙克托普Munktorp = BMunktorp蒙克托普
	f.蒙克托普Munktorp.SetParent(f)
	
	f.努尔贝里Norberg = BNorberg努尔贝里
	f.努尔贝里Norberg.SetParent(f)
	
	f.欣斯卡特贝里Skinnskatteberg = BSkinnskatteberg欣斯卡特贝里
	f.欣斯卡特贝里Skinnskatteberg.SetParent(f)
	
	f.韦斯特罗斯Vasteras = BVasteras韦斯特罗斯
	f.韦斯特罗斯Vasteras.SetParent(f)
	
}
