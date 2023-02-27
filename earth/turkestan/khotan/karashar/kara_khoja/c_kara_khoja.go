package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kara_khojaCounty interface {
    feud.County
    BAstana阿斯塔那() 	feud.Barony
    BBezeklik柏孜克里克() 	feud.Barony
    BHuhu狐胡() 	feud.Barony
    BJiaohe交河() 	feud.Barony
    BKara_khoja哈剌火州() 	feud.Barony
    BShangwangguo山王国() 	feud.Barony
    BTurfan吐鲁番() 	feud.Barony
}

type 哈剌火州Kara_khojaCounty struct {
	feud.BaseCounty
	阿斯塔那Astana 	feud.Barony
	柏孜克里克Bezeklik 	feud.Barony
	狐胡Huhu 	feud.Barony
	交河Jiaohe 	feud.Barony
	哈剌火州Kara_khoja 	feud.Barony
	山王国Shangwangguo 	feud.Barony
	吐鲁番Turfan 	feud.Barony
}

func (c *哈剌火州Kara_khojaCounty) BAstana阿斯塔那() feud.Barony {
	return c.阿斯塔那Astana
}
    
func (c *哈剌火州Kara_khojaCounty) BBezeklik柏孜克里克() feud.Barony {
	return c.柏孜克里克Bezeklik
}
    
func (c *哈剌火州Kara_khojaCounty) BHuhu狐胡() feud.Barony {
	return c.狐胡Huhu
}
    
func (c *哈剌火州Kara_khojaCounty) BJiaohe交河() feud.Barony {
	return c.交河Jiaohe
}
    
func (c *哈剌火州Kara_khojaCounty) BKara_khoja哈剌火州() feud.Barony {
	return c.哈剌火州Kara_khoja
}
    
func (c *哈剌火州Kara_khojaCounty) BShangwangguo山王国() feud.Barony {
	return c.山王国Shangwangguo
}
    
func (c *哈剌火州Kara_khojaCounty) BTurfan吐鲁番() feud.Barony {
	return c.吐鲁番Turfan
}
    
var CKara_khoja哈剌火州 Kara_khojaCounty = &哈剌火州Kara_khojaCounty{}

func init() {
	f := CKara_khoja哈剌火州.(*哈剌火州Kara_khojaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1446",
		Title:     "kara_khoja",
		TitleName: "哈剌火州",
		TitleCode: "c_kara_khoja",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯塔那Astana = BAstana阿斯塔那
	f.阿斯塔那Astana.SetParent(f)
	
	f.柏孜克里克Bezeklik = BBezeklik柏孜克里克
	f.柏孜克里克Bezeklik.SetParent(f)
	
	f.狐胡Huhu = BHuhu狐胡
	f.狐胡Huhu.SetParent(f)
	
	f.交河Jiaohe = BJiaohe交河
	f.交河Jiaohe.SetParent(f)
	
	f.哈剌火州Kara_khoja = BKara_khoja哈剌火州
	f.哈剌火州Kara_khoja.SetParent(f)
	
	f.山王国Shangwangguo = BShangwangguo山王国
	f.山王国Shangwangguo.SetParent(f)
	
	f.吐鲁番Turfan = BTurfan吐鲁番
	f.吐鲁番Turfan.SetParent(f)
	
}
