package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WagadougouCounty interface {
    feud.County
    BBalangoka巴朗戈卡() 	feud.Barony
    BBoboye博博耶() 	feud.Barony
    BDiba迪巴() 	feud.Barony
    BGursi古尔西() 	feud.Barony
    BNiamalo尼亚马洛() 	feud.Barony
    BPo波() 	feud.Barony
    BWagadougou瓦加杜古() 	feud.Barony
}

type 瓦加杜古WagadougouCounty struct {
	feud.BaseCounty
	巴朗戈卡Balangoka 	feud.Barony
	博博耶Boboye 	feud.Barony
	迪巴Diba 	feud.Barony
	古尔西Gursi 	feud.Barony
	尼亚马洛Niamalo 	feud.Barony
	波Po 	feud.Barony
	瓦加杜古Wagadougou 	feud.Barony
}

func (c *瓦加杜古WagadougouCounty) BBalangoka巴朗戈卡() feud.Barony {
	return c.巴朗戈卡Balangoka
}
    
func (c *瓦加杜古WagadougouCounty) BBoboye博博耶() feud.Barony {
	return c.博博耶Boboye
}
    
func (c *瓦加杜古WagadougouCounty) BDiba迪巴() feud.Barony {
	return c.迪巴Diba
}
    
func (c *瓦加杜古WagadougouCounty) BGursi古尔西() feud.Barony {
	return c.古尔西Gursi
}
    
func (c *瓦加杜古WagadougouCounty) BNiamalo尼亚马洛() feud.Barony {
	return c.尼亚马洛Niamalo
}
    
func (c *瓦加杜古WagadougouCounty) BPo波() feud.Barony {
	return c.波Po
}
    
func (c *瓦加杜古WagadougouCounty) BWagadougou瓦加杜古() feud.Barony {
	return c.瓦加杜古Wagadougou
}
    
var CWagadougou瓦加杜古 WagadougouCounty = &瓦加杜古WagadougouCounty{}

func init() {
	f := CWagadougou瓦加杜古.(*瓦加杜古WagadougouCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1759",
		Title:     "wagadougou",
		TitleName: "瓦加杜古",
		TitleCode: "c_wagadougou",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴朗戈卡Balangoka = BBalangoka巴朗戈卡
	f.巴朗戈卡Balangoka.SetParent(f)
	
	f.博博耶Boboye = BBoboye博博耶
	f.博博耶Boboye.SetParent(f)
	
	f.迪巴Diba = BDiba迪巴
	f.迪巴Diba.SetParent(f)
	
	f.古尔西Gursi = BGursi古尔西
	f.古尔西Gursi.SetParent(f)
	
	f.尼亚马洛Niamalo = BNiamalo尼亚马洛
	f.尼亚马洛Niamalo.SetParent(f)
	
	f.波Po = BPo波
	f.波Po.SetParent(f)
	
	f.瓦加杜古Wagadougou = BWagadougou瓦加杜古
	f.瓦加杜古Wagadougou.SetParent(f)
	
}
