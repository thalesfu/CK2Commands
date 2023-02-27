package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Lower_donCounty interface {
    feud.County
    BLatonovo拉托诺沃() 	feud.Barony
    BMarfinka马尔芬卡() 	feud.Barony
    BMatveevkurgan马特韦耶夫_库尔干() 	feud.Barony
    BNovoazovsk诺沃亚佐夫斯克() 	feud.Barony
    BRyasnoye里亚斯诺耶() 	feud.Barony
    BSkorokhod斯科罗霍德() 	feud.Barony
    BTaganrog塔甘罗格() 	feud.Barony
    BVesely韦谢雷() 	feud.Barony
}

type 下顿河Lower_donCounty struct {
	feud.BaseCounty
	拉托诺沃Latonovo 	feud.Barony
	马尔芬卡Marfinka 	feud.Barony
	马特韦耶夫_库尔干Matveevkurgan 	feud.Barony
	诺沃亚佐夫斯克Novoazovsk 	feud.Barony
	里亚斯诺耶Ryasnoye 	feud.Barony
	斯科罗霍德Skorokhod 	feud.Barony
	塔甘罗格Taganrog 	feud.Barony
	韦谢雷Vesely 	feud.Barony
}

func (c *下顿河Lower_donCounty) BLatonovo拉托诺沃() feud.Barony {
	return c.拉托诺沃Latonovo
}
    
func (c *下顿河Lower_donCounty) BMarfinka马尔芬卡() feud.Barony {
	return c.马尔芬卡Marfinka
}
    
func (c *下顿河Lower_donCounty) BMatveevkurgan马特韦耶夫_库尔干() feud.Barony {
	return c.马特韦耶夫_库尔干Matveevkurgan
}
    
func (c *下顿河Lower_donCounty) BNovoazovsk诺沃亚佐夫斯克() feud.Barony {
	return c.诺沃亚佐夫斯克Novoazovsk
}
    
func (c *下顿河Lower_donCounty) BRyasnoye里亚斯诺耶() feud.Barony {
	return c.里亚斯诺耶Ryasnoye
}
    
func (c *下顿河Lower_donCounty) BSkorokhod斯科罗霍德() feud.Barony {
	return c.斯科罗霍德Skorokhod
}
    
func (c *下顿河Lower_donCounty) BTaganrog塔甘罗格() feud.Barony {
	return c.塔甘罗格Taganrog
}
    
func (c *下顿河Lower_donCounty) BVesely韦谢雷() feud.Barony {
	return c.韦谢雷Vesely
}
    
var CLower_don下顿河 Lower_donCounty = &下顿河Lower_donCounty{}

func init() {
	f := CLower_don下顿河.(*下顿河Lower_donCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "563",
		Title:     "lower_don",
		TitleName: "下顿河",
		TitleCode: "c_lower_don",
		Baronies:  map[string]feud.Barony{},
	}

	f.拉托诺沃Latonovo = BLatonovo拉托诺沃
	f.拉托诺沃Latonovo.SetParent(f)
	
	f.马尔芬卡Marfinka = BMarfinka马尔芬卡
	f.马尔芬卡Marfinka.SetParent(f)
	
	f.马特韦耶夫_库尔干Matveevkurgan = BMatveevkurgan马特韦耶夫_库尔干
	f.马特韦耶夫_库尔干Matveevkurgan.SetParent(f)
	
	f.诺沃亚佐夫斯克Novoazovsk = BNovoazovsk诺沃亚佐夫斯克
	f.诺沃亚佐夫斯克Novoazovsk.SetParent(f)
	
	f.里亚斯诺耶Ryasnoye = BRyasnoye里亚斯诺耶
	f.里亚斯诺耶Ryasnoye.SetParent(f)
	
	f.斯科罗霍德Skorokhod = BSkorokhod斯科罗霍德
	f.斯科罗霍德Skorokhod.SetParent(f)
	
	f.塔甘罗格Taganrog = BTaganrog塔甘罗格
	f.塔甘罗格Taganrog.SetParent(f)
	
	f.韦谢雷Vesely = BVesely韦谢雷
	f.韦谢雷Vesely.SetParent(f)
	
}
