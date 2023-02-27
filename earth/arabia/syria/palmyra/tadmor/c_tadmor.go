package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TadmorCounty interface {
    feud.County
    BAlqaim加伊姆() 	feud.Barony
    BAlqunjah古里亚() 	feud.Barony
    BAsharah阿沙拉() 	feud.Barony
    BAssusah苏塞() 	feud.Barony
    BBukamal布凯马勒() 	feud.Barony
    BHusaiba胡塞巴() 	feud.Barony
    BSalhiyah赛希亚() 	feud.Barony
    BSubaykhan苏拜汉() 	feud.Barony
}

type 苏赫奈TadmorCounty struct {
	feud.BaseCounty
	加伊姆Alqaim 	feud.Barony
	古里亚Alqunjah 	feud.Barony
	阿沙拉Asharah 	feud.Barony
	苏塞Assusah 	feud.Barony
	布凯马勒Bukamal 	feud.Barony
	胡塞巴Husaiba 	feud.Barony
	赛希亚Salhiyah 	feud.Barony
	苏拜汉Subaykhan 	feud.Barony
}

func (c *苏赫奈TadmorCounty) BAlqaim加伊姆() feud.Barony {
	return c.加伊姆Alqaim
}
    
func (c *苏赫奈TadmorCounty) BAlqunjah古里亚() feud.Barony {
	return c.古里亚Alqunjah
}
    
func (c *苏赫奈TadmorCounty) BAsharah阿沙拉() feud.Barony {
	return c.阿沙拉Asharah
}
    
func (c *苏赫奈TadmorCounty) BAssusah苏塞() feud.Barony {
	return c.苏塞Assusah
}
    
func (c *苏赫奈TadmorCounty) BBukamal布凯马勒() feud.Barony {
	return c.布凯马勒Bukamal
}
    
func (c *苏赫奈TadmorCounty) BHusaiba胡塞巴() feud.Barony {
	return c.胡塞巴Husaiba
}
    
func (c *苏赫奈TadmorCounty) BSalhiyah赛希亚() feud.Barony {
	return c.赛希亚Salhiyah
}
    
func (c *苏赫奈TadmorCounty) BSubaykhan苏拜汉() feud.Barony {
	return c.苏拜汉Subaykhan
}
    
var CTadmor苏赫奈 TadmorCounty = &苏赫奈TadmorCounty{}

func init() {
	f := CTadmor苏赫奈.(*苏赫奈TadmorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "729",
		Title:     "tadmor",
		TitleName: "苏赫奈",
		TitleCode: "c_tadmor",
		Baronies:  map[string]feud.Barony{},
	}

	f.加伊姆Alqaim = BAlqaim加伊姆
	f.加伊姆Alqaim.SetParent(f)
	
	f.古里亚Alqunjah = BAlqunjah古里亚
	f.古里亚Alqunjah.SetParent(f)
	
	f.阿沙拉Asharah = BAsharah阿沙拉
	f.阿沙拉Asharah.SetParent(f)
	
	f.苏塞Assusah = BAssusah苏塞
	f.苏塞Assusah.SetParent(f)
	
	f.布凯马勒Bukamal = BBukamal布凯马勒
	f.布凯马勒Bukamal.SetParent(f)
	
	f.胡塞巴Husaiba = BHusaiba胡塞巴
	f.胡塞巴Husaiba.SetParent(f)
	
	f.赛希亚Salhiyah = BSalhiyah赛希亚
	f.赛希亚Salhiyah.SetParent(f)
	
	f.苏拜汉Subaykhan = BSubaykhan苏拜汉
	f.苏拜汉Subaykhan.SetParent(f)
	
}
