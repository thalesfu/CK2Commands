package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlcantaraCounty interface {
    feud.County
    BAlcantara阿尔坎塔拉() 	feud.Barony
    BBrozas布罗萨斯() 	feud.Barony
    BCeclavin塞克拉温() 	feud.Barony
    BCoria科里亚() 	feud.Barony
    BLamata拉马塔() 	feud.Barony
    BLasnavasdelmadrono拉斯纳瓦斯德尔马德罗尼奥() 	feud.Barony
    BMoraleja莫拉莱哈() 	feud.Barony
    BRacharachel拉查拉切尔() 	feud.Barony
}

type 阿尔坎塔拉AlcantaraCounty struct {
	feud.BaseCounty
	阿尔坎塔拉Alcantara 	feud.Barony
	布罗萨斯Brozas 	feud.Barony
	塞克拉温Ceclavin 	feud.Barony
	科里亚Coria 	feud.Barony
	拉马塔Lamata 	feud.Barony
	拉斯纳瓦斯德尔马德罗尼奥Lasnavasdelmadrono 	feud.Barony
	莫拉莱哈Moraleja 	feud.Barony
	拉查拉切尔Racharachel 	feud.Barony
}

func (c *阿尔坎塔拉AlcantaraCounty) BAlcantara阿尔坎塔拉() feud.Barony {
	return c.阿尔坎塔拉Alcantara
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BBrozas布罗萨斯() feud.Barony {
	return c.布罗萨斯Brozas
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BCeclavin塞克拉温() feud.Barony {
	return c.塞克拉温Ceclavin
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BCoria科里亚() feud.Barony {
	return c.科里亚Coria
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BLamata拉马塔() feud.Barony {
	return c.拉马塔Lamata
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BLasnavasdelmadrono拉斯纳瓦斯德尔马德罗尼奥() feud.Barony {
	return c.拉斯纳瓦斯德尔马德罗尼奥Lasnavasdelmadrono
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BMoraleja莫拉莱哈() feud.Barony {
	return c.莫拉莱哈Moraleja
}
    
func (c *阿尔坎塔拉AlcantaraCounty) BRacharachel拉查拉切尔() feud.Barony {
	return c.拉查拉切尔Racharachel
}
    
var CAlcantara阿尔坎塔拉 AlcantaraCounty = &阿尔坎塔拉AlcantaraCounty{}

func init() {
	f := CAlcantara阿尔坎塔拉.(*阿尔坎塔拉AlcantaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "193",
		Title:     "alcantara",
		TitleName: "阿尔坎塔拉",
		TitleCode: "c_alcantara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔坎塔拉Alcantara = BAlcantara阿尔坎塔拉
	f.阿尔坎塔拉Alcantara.SetParent(f)
	
	f.布罗萨斯Brozas = BBrozas布罗萨斯
	f.布罗萨斯Brozas.SetParent(f)
	
	f.塞克拉温Ceclavin = BCeclavin塞克拉温
	f.塞克拉温Ceclavin.SetParent(f)
	
	f.科里亚Coria = BCoria科里亚
	f.科里亚Coria.SetParent(f)
	
	f.拉马塔Lamata = BLamata拉马塔
	f.拉马塔Lamata.SetParent(f)
	
	f.拉斯纳瓦斯德尔马德罗尼奥Lasnavasdelmadrono = BLasnavasdelmadrono拉斯纳瓦斯德尔马德罗尼奥
	f.拉斯纳瓦斯德尔马德罗尼奥Lasnavasdelmadrono.SetParent(f)
	
	f.莫拉莱哈Moraleja = BMoraleja莫拉莱哈
	f.莫拉莱哈Moraleja.SetParent(f)
	
	f.拉查拉切尔Racharachel = BRacharachel拉查拉切尔
	f.拉查拉切尔Racharachel.SetParent(f)
	
}
