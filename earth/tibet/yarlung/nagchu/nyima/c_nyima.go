package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NyimaCounty interface {
    feud.County
    BChugtsodropo秋措卓波() 	feud.Barony
    BJialong加龙() 	feud.Barony
    BJiwa吉瓦() 	feud.Barony
    BKhyungdzong琼宗() 	feud.Barony
    BLaiduo来多() 	feud.Barony
    BNyima尼玛() 	feud.Barony
    BTanglai塘来() 	feud.Barony
}

type 尼玛NyimaCounty struct {
	feud.BaseCounty
	秋措卓波Chugtsodropo 	feud.Barony
	加龙Jialong 	feud.Barony
	吉瓦Jiwa 	feud.Barony
	琼宗Khyungdzong 	feud.Barony
	来多Laiduo 	feud.Barony
	尼玛Nyima 	feud.Barony
	塘来Tanglai 	feud.Barony
}

func (c *尼玛NyimaCounty) BChugtsodropo秋措卓波() feud.Barony {
	return c.秋措卓波Chugtsodropo
}
    
func (c *尼玛NyimaCounty) BJialong加龙() feud.Barony {
	return c.加龙Jialong
}
    
func (c *尼玛NyimaCounty) BJiwa吉瓦() feud.Barony {
	return c.吉瓦Jiwa
}
    
func (c *尼玛NyimaCounty) BKhyungdzong琼宗() feud.Barony {
	return c.琼宗Khyungdzong
}
    
func (c *尼玛NyimaCounty) BLaiduo来多() feud.Barony {
	return c.来多Laiduo
}
    
func (c *尼玛NyimaCounty) BNyima尼玛() feud.Barony {
	return c.尼玛Nyima
}
    
func (c *尼玛NyimaCounty) BTanglai塘来() feud.Barony {
	return c.塘来Tanglai
}
    
var CNyima尼玛 NyimaCounty = &尼玛NyimaCounty{}

func init() {
	f := CNyima尼玛.(*尼玛NyimaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1507",
		Title:     "nyima",
		TitleName: "尼玛",
		TitleCode: "c_nyima",
		Baronies:  map[string]feud.Barony{},
	}

	f.秋措卓波Chugtsodropo = BChugtsodropo秋措卓波
	f.秋措卓波Chugtsodropo.SetParent(f)
	
	f.加龙Jialong = BJialong加龙
	f.加龙Jialong.SetParent(f)
	
	f.吉瓦Jiwa = BJiwa吉瓦
	f.吉瓦Jiwa.SetParent(f)
	
	f.琼宗Khyungdzong = BKhyungdzong琼宗
	f.琼宗Khyungdzong.SetParent(f)
	
	f.来多Laiduo = BLaiduo来多
	f.来多Laiduo.SetParent(f)
	
	f.尼玛Nyima = BNyima尼玛
	f.尼玛Nyima.SetParent(f)
	
	f.塘来Tanglai = BTanglai塘来
	f.塘来Tanglai.SetParent(f)
	
}
