package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TuulCounty interface {
    feud.County
    BArgalant阿尔嘎朗特() 	feud.Barony
    BGachuurt嘎丘尔特() 	feud.Barony
    BKhustain呼斯坦() 	feud.Barony
    BLun_tuul隆() 	feud.Barony
    BNalaikh纳来哈() 	feud.Barony
    BTerelj特勒尔吉() 	feud.Barony
    BTuul独乐() 	feud.Barony
}

type 独乐TuulCounty struct {
	feud.BaseCounty
	阿尔嘎朗特Argalant 	feud.Barony
	嘎丘尔特Gachuurt 	feud.Barony
	呼斯坦Khustain 	feud.Barony
	隆Lun_tuul 	feud.Barony
	纳来哈Nalaikh 	feud.Barony
	特勒尔吉Terelj 	feud.Barony
	独乐Tuul 	feud.Barony
}

func (c *独乐TuulCounty) BArgalant阿尔嘎朗特() feud.Barony {
	return c.阿尔嘎朗特Argalant
}
    
func (c *独乐TuulCounty) BGachuurt嘎丘尔特() feud.Barony {
	return c.嘎丘尔特Gachuurt
}
    
func (c *独乐TuulCounty) BKhustain呼斯坦() feud.Barony {
	return c.呼斯坦Khustain
}
    
func (c *独乐TuulCounty) BLun_tuul隆() feud.Barony {
	return c.隆Lun_tuul
}
    
func (c *独乐TuulCounty) BNalaikh纳来哈() feud.Barony {
	return c.纳来哈Nalaikh
}
    
func (c *独乐TuulCounty) BTerelj特勒尔吉() feud.Barony {
	return c.特勒尔吉Terelj
}
    
func (c *独乐TuulCounty) BTuul独乐() feud.Barony {
	return c.独乐Tuul
}
    
var CTuul独乐 TuulCounty = &独乐TuulCounty{}

func init() {
	f := CTuul独乐.(*独乐TuulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1925",
		Title:     "tuul",
		TitleName: "独乐",
		TitleCode: "c_tuul",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔嘎朗特Argalant = BArgalant阿尔嘎朗特
	f.阿尔嘎朗特Argalant.SetParent(f)
	
	f.嘎丘尔特Gachuurt = BGachuurt嘎丘尔特
	f.嘎丘尔特Gachuurt.SetParent(f)
	
	f.呼斯坦Khustain = BKhustain呼斯坦
	f.呼斯坦Khustain.SetParent(f)
	
	f.隆Lun_tuul = BLun_tuul隆
	f.隆Lun_tuul.SetParent(f)
	
	f.纳来哈Nalaikh = BNalaikh纳来哈
	f.纳来哈Nalaikh.SetParent(f)
	
	f.特勒尔吉Terelj = BTerelj特勒尔吉
	f.特勒尔吉Terelj.SetParent(f)
	
	f.独乐Tuul = BTuul独乐
	f.独乐Tuul.SetParent(f)
	
}
