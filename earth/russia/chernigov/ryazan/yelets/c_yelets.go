package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YeletsCounty interface {
    feud.County
    BDolgorukovo多尔戈鲁科沃() 	feud.Barony
    BKhlevnoye赫列夫诺耶() 	feud.Barony
    BSemiuki谢米卢基() 	feud.Barony
    BTerbuny捷尔布内() 	feud.Barony
    BVoronej沃罗涅日() 	feud.Barony
    BYelets叶列茨() 	feud.Barony
    BZadonsk扎顿斯克() 	feud.Barony
}

type 叶列茨YeletsCounty struct {
	feud.BaseCounty
	多尔戈鲁科沃Dolgorukovo 	feud.Barony
	赫列夫诺耶Khlevnoye 	feud.Barony
	谢米卢基Semiuki 	feud.Barony
	捷尔布内Terbuny 	feud.Barony
	沃罗涅日Voronej 	feud.Barony
	叶列茨Yelets 	feud.Barony
	扎顿斯克Zadonsk 	feud.Barony
}

func (c *叶列茨YeletsCounty) BDolgorukovo多尔戈鲁科沃() feud.Barony {
	return c.多尔戈鲁科沃Dolgorukovo
}
    
func (c *叶列茨YeletsCounty) BKhlevnoye赫列夫诺耶() feud.Barony {
	return c.赫列夫诺耶Khlevnoye
}
    
func (c *叶列茨YeletsCounty) BSemiuki谢米卢基() feud.Barony {
	return c.谢米卢基Semiuki
}
    
func (c *叶列茨YeletsCounty) BTerbuny捷尔布内() feud.Barony {
	return c.捷尔布内Terbuny
}
    
func (c *叶列茨YeletsCounty) BVoronej沃罗涅日() feud.Barony {
	return c.沃罗涅日Voronej
}
    
func (c *叶列茨YeletsCounty) BYelets叶列茨() feud.Barony {
	return c.叶列茨Yelets
}
    
func (c *叶列茨YeletsCounty) BZadonsk扎顿斯克() feud.Barony {
	return c.扎顿斯克Zadonsk
}
    
var CYelets叶列茨 YeletsCounty = &叶列茨YeletsCounty{}

func init() {
	f := CYelets叶列茨.(*叶列茨YeletsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1665",
		Title:     "yelets",
		TitleName: "叶列茨",
		TitleCode: "c_yelets",
		Baronies:  map[string]feud.Barony{},
	}

	f.多尔戈鲁科沃Dolgorukovo = BDolgorukovo多尔戈鲁科沃
	f.多尔戈鲁科沃Dolgorukovo.SetParent(f)
	
	f.赫列夫诺耶Khlevnoye = BKhlevnoye赫列夫诺耶
	f.赫列夫诺耶Khlevnoye.SetParent(f)
	
	f.谢米卢基Semiuki = BSemiuki谢米卢基
	f.谢米卢基Semiuki.SetParent(f)
	
	f.捷尔布内Terbuny = BTerbuny捷尔布内
	f.捷尔布内Terbuny.SetParent(f)
	
	f.沃罗涅日Voronej = BVoronej沃罗涅日
	f.沃罗涅日Voronej.SetParent(f)
	
	f.叶列茨Yelets = BYelets叶列茨
	f.叶列茨Yelets.SetParent(f)
	
	f.扎顿斯克Zadonsk = BZadonsk扎顿斯克
	f.扎顿斯克Zadonsk.SetParent(f)
	
}
