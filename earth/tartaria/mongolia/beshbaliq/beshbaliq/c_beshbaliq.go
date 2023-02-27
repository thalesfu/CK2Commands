package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeshbaliqCounty interface {
    feud.County
    BBashqorghan巴什库尔干() 	feud.Barony
    BBeshbalik别失八里() 	feud.Barony
    BGucheng古城() 	feud.Barony
    BJiangjunmiao将军庙() 	feud.Barony
    BLishan骊山() 	feud.Barony
    BLiuwanxiang柳湾乡() 	feud.Barony
    BPinghu平湖() 	feud.Barony
}

type 别失八里BeshbaliqCounty struct {
	feud.BaseCounty
	巴什库尔干Bashqorghan 	feud.Barony
	别失八里Beshbalik 	feud.Barony
	古城Gucheng 	feud.Barony
	将军庙Jiangjunmiao 	feud.Barony
	骊山Lishan 	feud.Barony
	柳湾乡Liuwanxiang 	feud.Barony
	平湖Pinghu 	feud.Barony
}

func (c *别失八里BeshbaliqCounty) BBashqorghan巴什库尔干() feud.Barony {
	return c.巴什库尔干Bashqorghan
}
    
func (c *别失八里BeshbaliqCounty) BBeshbalik别失八里() feud.Barony {
	return c.别失八里Beshbalik
}
    
func (c *别失八里BeshbaliqCounty) BGucheng古城() feud.Barony {
	return c.古城Gucheng
}
    
func (c *别失八里BeshbaliqCounty) BJiangjunmiao将军庙() feud.Barony {
	return c.将军庙Jiangjunmiao
}
    
func (c *别失八里BeshbaliqCounty) BLishan骊山() feud.Barony {
	return c.骊山Lishan
}
    
func (c *别失八里BeshbaliqCounty) BLiuwanxiang柳湾乡() feud.Barony {
	return c.柳湾乡Liuwanxiang
}
    
func (c *别失八里BeshbaliqCounty) BPinghu平湖() feud.Barony {
	return c.平湖Pinghu
}
    
var CBeshbaliq别失八里 BeshbaliqCounty = &别失八里BeshbaliqCounty{}

func init() {
	f := CBeshbaliq别失八里.(*别失八里BeshbaliqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1452",
		Title:     "beshbaliq",
		TitleName: "别失八里",
		TitleCode: "c_beshbaliq",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴什库尔干Bashqorghan = BBashqorghan巴什库尔干
	f.巴什库尔干Bashqorghan.SetParent(f)
	
	f.别失八里Beshbalik = BBeshbalik别失八里
	f.别失八里Beshbalik.SetParent(f)
	
	f.古城Gucheng = BGucheng古城
	f.古城Gucheng.SetParent(f)
	
	f.将军庙Jiangjunmiao = BJiangjunmiao将军庙
	f.将军庙Jiangjunmiao.SetParent(f)
	
	f.骊山Lishan = BLishan骊山
	f.骊山Lishan.SetParent(f)
	
	f.柳湾乡Liuwanxiang = BLiuwanxiang柳湾乡
	f.柳湾乡Liuwanxiang.SetParent(f)
	
	f.平湖Pinghu = BPinghu平湖
	f.平湖Pinghu.SetParent(f)
	
}
