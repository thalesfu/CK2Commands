package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GondarCounty interface {
    feud.County
    BAmbassel阿姆巴瑟() 	feud.Barony
    BBahirdar巴赫达尔() 	feud.Barony
    BDembiya德姆比亚() 	feud.Barony
    BFasilghebbi法西尔盖比() 	feud.Barony
    BGondar贡德尔() 	feud.Barony
    BMagdala马格达拉() 	feud.Barony
    BRoha罗哈() 	feud.Barony
    BTanaqirqos塔纳奇果斯() 	feud.Barony
}

type 贡德尔GondarCounty struct {
	feud.BaseCounty
	阿姆巴瑟Ambassel 	feud.Barony
	巴赫达尔Bahirdar 	feud.Barony
	德姆比亚Dembiya 	feud.Barony
	法西尔盖比Fasilghebbi 	feud.Barony
	贡德尔Gondar 	feud.Barony
	马格达拉Magdala 	feud.Barony
	罗哈Roha 	feud.Barony
	塔纳奇果斯Tanaqirqos 	feud.Barony
}

func (c *贡德尔GondarCounty) BAmbassel阿姆巴瑟() feud.Barony {
	return c.阿姆巴瑟Ambassel
}
    
func (c *贡德尔GondarCounty) BBahirdar巴赫达尔() feud.Barony {
	return c.巴赫达尔Bahirdar
}
    
func (c *贡德尔GondarCounty) BDembiya德姆比亚() feud.Barony {
	return c.德姆比亚Dembiya
}
    
func (c *贡德尔GondarCounty) BFasilghebbi法西尔盖比() feud.Barony {
	return c.法西尔盖比Fasilghebbi
}
    
func (c *贡德尔GondarCounty) BGondar贡德尔() feud.Barony {
	return c.贡德尔Gondar
}
    
func (c *贡德尔GondarCounty) BMagdala马格达拉() feud.Barony {
	return c.马格达拉Magdala
}
    
func (c *贡德尔GondarCounty) BRoha罗哈() feud.Barony {
	return c.罗哈Roha
}
    
func (c *贡德尔GondarCounty) BTanaqirqos塔纳奇果斯() feud.Barony {
	return c.塔纳奇果斯Tanaqirqos
}
    
var CGondar贡德尔 GondarCounty = &贡德尔GondarCounty{}

func init() {
	f := CGondar贡德尔.(*贡德尔GondarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "883",
		Title:     "gondar",
		TitleName: "贡德尔",
		TitleCode: "c_gondar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆巴瑟Ambassel = BAmbassel阿姆巴瑟
	f.阿姆巴瑟Ambassel.SetParent(f)
	
	f.巴赫达尔Bahirdar = BBahirdar巴赫达尔
	f.巴赫达尔Bahirdar.SetParent(f)
	
	f.德姆比亚Dembiya = BDembiya德姆比亚
	f.德姆比亚Dembiya.SetParent(f)
	
	f.法西尔盖比Fasilghebbi = BFasilghebbi法西尔盖比
	f.法西尔盖比Fasilghebbi.SetParent(f)
	
	f.贡德尔Gondar = BGondar贡德尔
	f.贡德尔Gondar.SetParent(f)
	
	f.马格达拉Magdala = BMagdala马格达拉
	f.马格达拉Magdala.SetParent(f)
	
	f.罗哈Roha = BRoha罗哈
	f.罗哈Roha.SetParent(f)
	
	f.塔纳奇果斯Tanaqirqos = BTanaqirqos塔纳奇果斯
	f.塔纳奇果斯Tanaqirqos.SetParent(f)
	
}
