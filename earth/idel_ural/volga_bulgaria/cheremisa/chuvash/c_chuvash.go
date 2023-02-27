package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChuvashCounty interface {
    feud.County
    BAlatyr阿拉特里() 	feud.Barony
    BCheboksary切博克萨雷() 	feud.Barony
    BKozlovka科兹洛夫卡() 	feud.Barony
    BMakaryevo马卡里耶沃() 	feud.Barony
    BSundyr孙德里() 	feud.Barony
    BTsivilsk齐维利斯克() 	feud.Barony
    BVedasuvar维达苏瓦尔() 	feud.Barony
    BYadrin亚德林() 	feud.Barony
}

type 楚瓦什ChuvashCounty struct {
	feud.BaseCounty
	阿拉特里Alatyr 	feud.Barony
	切博克萨雷Cheboksary 	feud.Barony
	科兹洛夫卡Kozlovka 	feud.Barony
	马卡里耶沃Makaryevo 	feud.Barony
	孙德里Sundyr 	feud.Barony
	齐维利斯克Tsivilsk 	feud.Barony
	维达苏瓦尔Vedasuvar 	feud.Barony
	亚德林Yadrin 	feud.Barony
}

func (c *楚瓦什ChuvashCounty) BAlatyr阿拉特里() feud.Barony {
	return c.阿拉特里Alatyr
}
    
func (c *楚瓦什ChuvashCounty) BCheboksary切博克萨雷() feud.Barony {
	return c.切博克萨雷Cheboksary
}
    
func (c *楚瓦什ChuvashCounty) BKozlovka科兹洛夫卡() feud.Barony {
	return c.科兹洛夫卡Kozlovka
}
    
func (c *楚瓦什ChuvashCounty) BMakaryevo马卡里耶沃() feud.Barony {
	return c.马卡里耶沃Makaryevo
}
    
func (c *楚瓦什ChuvashCounty) BSundyr孙德里() feud.Barony {
	return c.孙德里Sundyr
}
    
func (c *楚瓦什ChuvashCounty) BTsivilsk齐维利斯克() feud.Barony {
	return c.齐维利斯克Tsivilsk
}
    
func (c *楚瓦什ChuvashCounty) BVedasuvar维达苏瓦尔() feud.Barony {
	return c.维达苏瓦尔Vedasuvar
}
    
func (c *楚瓦什ChuvashCounty) BYadrin亚德林() feud.Barony {
	return c.亚德林Yadrin
}
    
var CChuvash楚瓦什 ChuvashCounty = &楚瓦什ChuvashCounty{}

func init() {
	f := CChuvash楚瓦什.(*楚瓦什ChuvashCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "590",
		Title:     "chuvash",
		TitleName: "楚瓦什",
		TitleCode: "c_chuvash",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉特里Alatyr = BAlatyr阿拉特里
	f.阿拉特里Alatyr.SetParent(f)
	
	f.切博克萨雷Cheboksary = BCheboksary切博克萨雷
	f.切博克萨雷Cheboksary.SetParent(f)
	
	f.科兹洛夫卡Kozlovka = BKozlovka科兹洛夫卡
	f.科兹洛夫卡Kozlovka.SetParent(f)
	
	f.马卡里耶沃Makaryevo = BMakaryevo马卡里耶沃
	f.马卡里耶沃Makaryevo.SetParent(f)
	
	f.孙德里Sundyr = BSundyr孙德里
	f.孙德里Sundyr.SetParent(f)
	
	f.齐维利斯克Tsivilsk = BTsivilsk齐维利斯克
	f.齐维利斯克Tsivilsk.SetParent(f)
	
	f.维达苏瓦尔Vedasuvar = BVedasuvar维达苏瓦尔
	f.维达苏瓦尔Vedasuvar.SetParent(f)
	
	f.亚德林Yadrin = BYadrin亚德林
	f.亚德林Yadrin.SetParent(f)
	
}
