package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalindiaCounty interface {
    feud.County
    BAngerburg安格堡() 	feud.Barony
    BGilgenburg吉尔根堡() 	feud.Barony
    BHohenstein霍亨施泰因() 	feud.Barony
    BNeidenburg奈登堡() 	feud.Barony
    BNikelshagen尼克尔斯哈根() 	feud.Barony
    BOsterode奥斯特罗德() 	feud.Barony
    BTreuburg特罗伊堡() 	feud.Barony
    BWielbark维尔巴克() 	feud.Barony
}

type 加林迪亚GalindiaCounty struct {
	feud.BaseCounty
	安格堡Angerburg 	feud.Barony
	吉尔根堡Gilgenburg 	feud.Barony
	霍亨施泰因Hohenstein 	feud.Barony
	奈登堡Neidenburg 	feud.Barony
	尼克尔斯哈根Nikelshagen 	feud.Barony
	奥斯特罗德Osterode 	feud.Barony
	特罗伊堡Treuburg 	feud.Barony
	维尔巴克Wielbark 	feud.Barony
}

func (c *加林迪亚GalindiaCounty) BAngerburg安格堡() feud.Barony {
	return c.安格堡Angerburg
}
    
func (c *加林迪亚GalindiaCounty) BGilgenburg吉尔根堡() feud.Barony {
	return c.吉尔根堡Gilgenburg
}
    
func (c *加林迪亚GalindiaCounty) BHohenstein霍亨施泰因() feud.Barony {
	return c.霍亨施泰因Hohenstein
}
    
func (c *加林迪亚GalindiaCounty) BNeidenburg奈登堡() feud.Barony {
	return c.奈登堡Neidenburg
}
    
func (c *加林迪亚GalindiaCounty) BNikelshagen尼克尔斯哈根() feud.Barony {
	return c.尼克尔斯哈根Nikelshagen
}
    
func (c *加林迪亚GalindiaCounty) BOsterode奥斯特罗德() feud.Barony {
	return c.奥斯特罗德Osterode
}
    
func (c *加林迪亚GalindiaCounty) BTreuburg特罗伊堡() feud.Barony {
	return c.特罗伊堡Treuburg
}
    
func (c *加林迪亚GalindiaCounty) BWielbark维尔巴克() feud.Barony {
	return c.维尔巴克Wielbark
}
    
var CGalindia加林迪亚 GalindiaCounty = &加林迪亚GalindiaCounty{}

func init() {
	f := CGalindia加林迪亚.(*加林迪亚GalindiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "427",
		Title:     "galindia",
		TitleName: "加林迪亚",
		TitleCode: "c_galindia",
		Baronies:  map[string]feud.Barony{},
	}

	f.安格堡Angerburg = BAngerburg安格堡
	f.安格堡Angerburg.SetParent(f)
	
	f.吉尔根堡Gilgenburg = BGilgenburg吉尔根堡
	f.吉尔根堡Gilgenburg.SetParent(f)
	
	f.霍亨施泰因Hohenstein = BHohenstein霍亨施泰因
	f.霍亨施泰因Hohenstein.SetParent(f)
	
	f.奈登堡Neidenburg = BNeidenburg奈登堡
	f.奈登堡Neidenburg.SetParent(f)
	
	f.尼克尔斯哈根Nikelshagen = BNikelshagen尼克尔斯哈根
	f.尼克尔斯哈根Nikelshagen.SetParent(f)
	
	f.奥斯特罗德Osterode = BOsterode奥斯特罗德
	f.奥斯特罗德Osterode.SetParent(f)
	
	f.特罗伊堡Treuburg = BTreuburg特罗伊堡
	f.特罗伊堡Treuburg.SetParent(f)
	
	f.维尔巴克Wielbark = BWielbark维尔巴克
	f.维尔巴克Wielbark.SetParent(f)
	
}
