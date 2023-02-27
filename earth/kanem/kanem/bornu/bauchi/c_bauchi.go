package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BauchiCounty interface {
    feud.County
    BBakinkuja巴金库贾() 	feud.Barony
    BBauchi包奇() 	feud.Barony
    BDanmadahun丹马达洪() 	feud.Barony
    BKadira卡迪拉() 	feud.Barony
    BKazaure卡扎乌雷() 	feud.Barony
    BRuwan_jema鲁万杰马() 	feud.Barony
    BZar扎尔() 	feud.Barony
}

type 包奇BauchiCounty struct {
	feud.BaseCounty
	巴金库贾Bakinkuja 	feud.Barony
	包奇Bauchi 	feud.Barony
	丹马达洪Danmadahun 	feud.Barony
	卡迪拉Kadira 	feud.Barony
	卡扎乌雷Kazaure 	feud.Barony
	鲁万杰马Ruwan_jema 	feud.Barony
	扎尔Zar 	feud.Barony
}

func (c *包奇BauchiCounty) BBakinkuja巴金库贾() feud.Barony {
	return c.巴金库贾Bakinkuja
}
    
func (c *包奇BauchiCounty) BBauchi包奇() feud.Barony {
	return c.包奇Bauchi
}
    
func (c *包奇BauchiCounty) BDanmadahun丹马达洪() feud.Barony {
	return c.丹马达洪Danmadahun
}
    
func (c *包奇BauchiCounty) BKadira卡迪拉() feud.Barony {
	return c.卡迪拉Kadira
}
    
func (c *包奇BauchiCounty) BKazaure卡扎乌雷() feud.Barony {
	return c.卡扎乌雷Kazaure
}
    
func (c *包奇BauchiCounty) BRuwan_jema鲁万杰马() feud.Barony {
	return c.鲁万杰马Ruwan_jema
}
    
func (c *包奇BauchiCounty) BZar扎尔() feud.Barony {
	return c.扎尔Zar
}
    
var CBauchi包奇 BauchiCounty = &包奇BauchiCounty{}

func init() {
	f := CBauchi包奇.(*包奇BauchiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1745",
		Title:     "bauchi",
		TitleName: "包奇",
		TitleCode: "c_bauchi",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴金库贾Bakinkuja = BBakinkuja巴金库贾
	f.巴金库贾Bakinkuja.SetParent(f)
	
	f.包奇Bauchi = BBauchi包奇
	f.包奇Bauchi.SetParent(f)
	
	f.丹马达洪Danmadahun = BDanmadahun丹马达洪
	f.丹马达洪Danmadahun.SetParent(f)
	
	f.卡迪拉Kadira = BKadira卡迪拉
	f.卡迪拉Kadira.SetParent(f)
	
	f.卡扎乌雷Kazaure = BKazaure卡扎乌雷
	f.卡扎乌雷Kazaure.SetParent(f)
	
	f.鲁万杰马Ruwan_jema = BRuwan_jema鲁万杰马
	f.鲁万杰马Ruwan_jema.SetParent(f)
	
	f.扎尔Zar = BZar扎尔
	f.扎尔Zar.SetParent(f)
	
}
