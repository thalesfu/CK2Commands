package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PurushapuraCounty interface {
    feud.County
    BAli_masjid阿里麦斯吉德() 	feud.Barony
    BBajaur巴焦尔() 	feud.Barony
    BMardan马尔丹() 	feud.Barony
    BNowshera瑙谢拉() 	feud.Barony
    BPurushapura布路沙布逻() 	feud.Barony
    BShah_ji_dheri沙_吉_谢里() 	feud.Barony
    BShergarh室罗姞利呬() 	feud.Barony
}

type 布路沙布逻PurushapuraCounty struct {
	feud.BaseCounty
	阿里麦斯吉德Ali_masjid 	feud.Barony
	巴焦尔Bajaur 	feud.Barony
	马尔丹Mardan 	feud.Barony
	瑙谢拉Nowshera 	feud.Barony
	布路沙布逻Purushapura 	feud.Barony
	沙_吉_谢里Shah_ji_dheri 	feud.Barony
	室罗姞利呬Shergarh 	feud.Barony
}

func (c *布路沙布逻PurushapuraCounty) BAli_masjid阿里麦斯吉德() feud.Barony {
	return c.阿里麦斯吉德Ali_masjid
}
    
func (c *布路沙布逻PurushapuraCounty) BBajaur巴焦尔() feud.Barony {
	return c.巴焦尔Bajaur
}
    
func (c *布路沙布逻PurushapuraCounty) BMardan马尔丹() feud.Barony {
	return c.马尔丹Mardan
}
    
func (c *布路沙布逻PurushapuraCounty) BNowshera瑙谢拉() feud.Barony {
	return c.瑙谢拉Nowshera
}
    
func (c *布路沙布逻PurushapuraCounty) BPurushapura布路沙布逻() feud.Barony {
	return c.布路沙布逻Purushapura
}
    
func (c *布路沙布逻PurushapuraCounty) BShah_ji_dheri沙_吉_谢里() feud.Barony {
	return c.沙_吉_谢里Shah_ji_dheri
}
    
func (c *布路沙布逻PurushapuraCounty) BShergarh室罗姞利呬() feud.Barony {
	return c.室罗姞利呬Shergarh
}
    
var CPurushapura布路沙布逻 PurushapuraCounty = &布路沙布逻PurushapuraCounty{}

func init() {
	f := CPurushapura布路沙布逻.(*布路沙布逻PurushapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1342",
		Title:     "purushapura",
		TitleName: "布路沙布逻",
		TitleCode: "c_purushapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿里麦斯吉德Ali_masjid = BAli_masjid阿里麦斯吉德
	f.阿里麦斯吉德Ali_masjid.SetParent(f)
	
	f.巴焦尔Bajaur = BBajaur巴焦尔
	f.巴焦尔Bajaur.SetParent(f)
	
	f.马尔丹Mardan = BMardan马尔丹
	f.马尔丹Mardan.SetParent(f)
	
	f.瑙谢拉Nowshera = BNowshera瑙谢拉
	f.瑙谢拉Nowshera.SetParent(f)
	
	f.布路沙布逻Purushapura = BPurushapura布路沙布逻
	f.布路沙布逻Purushapura.SetParent(f)
	
	f.沙_吉_谢里Shah_ji_dheri = BShah_ji_dheri沙_吉_谢里
	f.沙_吉_谢里Shah_ji_dheri.SetParent(f)
	
	f.室罗姞利呬Shergarh = BShergarh室罗姞利呬
	f.室罗姞利呬Shergarh.SetParent(f)
	
}
