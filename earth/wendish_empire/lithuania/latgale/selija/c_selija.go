package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SelijaCounty interface {
    feud.County
    BDignaja迪格纳亚() 	feud.Barony
    BDubenou杜贝纽() 	feud.Barony
    BLichsta利奇斯塔() 	feud.Barony
    BOzikel奥齐凯尔() 	feud.Barony
    BRackiska拉基什卡() 	feud.Barony
    BRurezeme库尔泽梅() 	feud.Barony
    BSelpils塞尔皮尔斯() 	feud.Barony
    BVarviai瓦尔维艾() 	feud.Barony
}

type 塞丽亚SelijaCounty struct {
	feud.BaseCounty
	迪格纳亚Dignaja 	feud.Barony
	杜贝纽Dubenou 	feud.Barony
	利奇斯塔Lichsta 	feud.Barony
	奥齐凯尔Ozikel 	feud.Barony
	拉基什卡Rackiska 	feud.Barony
	库尔泽梅Rurezeme 	feud.Barony
	塞尔皮尔斯Selpils 	feud.Barony
	瓦尔维艾Varviai 	feud.Barony
}

func (c *塞丽亚SelijaCounty) BDignaja迪格纳亚() feud.Barony {
	return c.迪格纳亚Dignaja
}
    
func (c *塞丽亚SelijaCounty) BDubenou杜贝纽() feud.Barony {
	return c.杜贝纽Dubenou
}
    
func (c *塞丽亚SelijaCounty) BLichsta利奇斯塔() feud.Barony {
	return c.利奇斯塔Lichsta
}
    
func (c *塞丽亚SelijaCounty) BOzikel奥齐凯尔() feud.Barony {
	return c.奥齐凯尔Ozikel
}
    
func (c *塞丽亚SelijaCounty) BRackiska拉基什卡() feud.Barony {
	return c.拉基什卡Rackiska
}
    
func (c *塞丽亚SelijaCounty) BRurezeme库尔泽梅() feud.Barony {
	return c.库尔泽梅Rurezeme
}
    
func (c *塞丽亚SelijaCounty) BSelpils塞尔皮尔斯() feud.Barony {
	return c.塞尔皮尔斯Selpils
}
    
func (c *塞丽亚SelijaCounty) BVarviai瓦尔维艾() feud.Barony {
	return c.瓦尔维艾Varviai
}
    
var CSelija塞丽亚 SelijaCounty = &塞丽亚SelijaCounty{}

func init() {
	f := CSelija塞丽亚.(*塞丽亚SelijaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1591",
		Title:     "selija",
		TitleName: "塞丽亚",
		TitleCode: "c_selija",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪格纳亚Dignaja = BDignaja迪格纳亚
	f.迪格纳亚Dignaja.SetParent(f)
	
	f.杜贝纽Dubenou = BDubenou杜贝纽
	f.杜贝纽Dubenou.SetParent(f)
	
	f.利奇斯塔Lichsta = BLichsta利奇斯塔
	f.利奇斯塔Lichsta.SetParent(f)
	
	f.奥齐凯尔Ozikel = BOzikel奥齐凯尔
	f.奥齐凯尔Ozikel.SetParent(f)
	
	f.拉基什卡Rackiska = BRackiska拉基什卡
	f.拉基什卡Rackiska.SetParent(f)
	
	f.库尔泽梅Rurezeme = BRurezeme库尔泽梅
	f.库尔泽梅Rurezeme.SetParent(f)
	
	f.塞尔皮尔斯Selpils = BSelpils塞尔皮尔斯
	f.塞尔皮尔斯Selpils.SetParent(f)
	
	f.瓦尔维艾Varviai = BVarviai瓦尔维艾
	f.瓦尔维艾Varviai.SetParent(f)
	
}
