package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarkelCounty interface {
    feud.County
    BBelayavezha白塔() 	feud.Barony
    BKazarki卡扎尔基() 	feud.Barony
    BKotelnikovo科捷利尼科沃() 	feud.Barony
    BNizhchir下奇尔() 	feud.Barony
    BSarkel萨克尔() 	feud.Barony
    BSemikarakorsk谢米卡拉科尔斯克() 	feud.Barony
    BTsimlyanskoye齐姆良斯科耶() 	feud.Barony
    BUstdonetskiy乌斯季_顿涅茨基() 	feud.Barony
}

type 萨克尔SarkelCounty struct {
	feud.BaseCounty
	白塔Belayavezha 	feud.Barony
	卡扎尔基Kazarki 	feud.Barony
	科捷利尼科沃Kotelnikovo 	feud.Barony
	下奇尔Nizhchir 	feud.Barony
	萨克尔Sarkel 	feud.Barony
	谢米卡拉科尔斯克Semikarakorsk 	feud.Barony
	齐姆良斯科耶Tsimlyanskoye 	feud.Barony
	乌斯季_顿涅茨基Ustdonetskiy 	feud.Barony
}

func (c *萨克尔SarkelCounty) BBelayavezha白塔() feud.Barony {
	return c.白塔Belayavezha
}
    
func (c *萨克尔SarkelCounty) BKazarki卡扎尔基() feud.Barony {
	return c.卡扎尔基Kazarki
}
    
func (c *萨克尔SarkelCounty) BKotelnikovo科捷利尼科沃() feud.Barony {
	return c.科捷利尼科沃Kotelnikovo
}
    
func (c *萨克尔SarkelCounty) BNizhchir下奇尔() feud.Barony {
	return c.下奇尔Nizhchir
}
    
func (c *萨克尔SarkelCounty) BSarkel萨克尔() feud.Barony {
	return c.萨克尔Sarkel
}
    
func (c *萨克尔SarkelCounty) BSemikarakorsk谢米卡拉科尔斯克() feud.Barony {
	return c.谢米卡拉科尔斯克Semikarakorsk
}
    
func (c *萨克尔SarkelCounty) BTsimlyanskoye齐姆良斯科耶() feud.Barony {
	return c.齐姆良斯科耶Tsimlyanskoye
}
    
func (c *萨克尔SarkelCounty) BUstdonetskiy乌斯季_顿涅茨基() feud.Barony {
	return c.乌斯季_顿涅茨基Ustdonetskiy
}
    
var CSarkel萨克尔 SarkelCounty = &萨克尔SarkelCounty{}

func init() {
	f := CSarkel萨克尔.(*萨克尔SarkelCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "594",
		Title:     "sarkel",
		TitleName: "萨克尔",
		TitleCode: "c_sarkel",
		Baronies:  map[string]feud.Barony{},
	}

	f.白塔Belayavezha = BBelayavezha白塔
	f.白塔Belayavezha.SetParent(f)
	
	f.卡扎尔基Kazarki = BKazarki卡扎尔基
	f.卡扎尔基Kazarki.SetParent(f)
	
	f.科捷利尼科沃Kotelnikovo = BKotelnikovo科捷利尼科沃
	f.科捷利尼科沃Kotelnikovo.SetParent(f)
	
	f.下奇尔Nizhchir = BNizhchir下奇尔
	f.下奇尔Nizhchir.SetParent(f)
	
	f.萨克尔Sarkel = BSarkel萨克尔
	f.萨克尔Sarkel.SetParent(f)
	
	f.谢米卡拉科尔斯克Semikarakorsk = BSemikarakorsk谢米卡拉科尔斯克
	f.谢米卡拉科尔斯克Semikarakorsk.SetParent(f)
	
	f.齐姆良斯科耶Tsimlyanskoye = BTsimlyanskoye齐姆良斯科耶
	f.齐姆良斯科耶Tsimlyanskoye.SetParent(f)
	
	f.乌斯季_顿涅茨基Ustdonetskiy = BUstdonetskiy乌斯季_顿涅茨基
	f.乌斯季_顿涅茨基Ustdonetskiy.SetParent(f)
	
}
