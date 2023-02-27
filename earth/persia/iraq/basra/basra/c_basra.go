package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BasraCounty interface {
    feud.County
    BArah阿拉() 	feud.Barony
    BAzzubayr祖拜尔() 	feud.Barony
    BBasra巴士拉() 	feud.Barony
    BKalaatderat卡拉特德拉特() 	feud.Barony
    BMohammera莫哈梅拉() 	feud.Barony
    BQurna喀尔那() 	feud.Barony
    BSukelsheyuhk苏克舍赫克() 	feud.Barony
    BUmmqasr乌姆盖萨尔() 	feud.Barony
}

type 巴士拉BasraCounty struct {
	feud.BaseCounty
	阿拉Arah 	feud.Barony
	祖拜尔Azzubayr 	feud.Barony
	巴士拉Basra 	feud.Barony
	卡拉特德拉特Kalaatderat 	feud.Barony
	莫哈梅拉Mohammera 	feud.Barony
	喀尔那Qurna 	feud.Barony
	苏克舍赫克Sukelsheyuhk 	feud.Barony
	乌姆盖萨尔Ummqasr 	feud.Barony
}

func (c *巴士拉BasraCounty) BArah阿拉() feud.Barony {
	return c.阿拉Arah
}
    
func (c *巴士拉BasraCounty) BAzzubayr祖拜尔() feud.Barony {
	return c.祖拜尔Azzubayr
}
    
func (c *巴士拉BasraCounty) BBasra巴士拉() feud.Barony {
	return c.巴士拉Basra
}
    
func (c *巴士拉BasraCounty) BKalaatderat卡拉特德拉特() feud.Barony {
	return c.卡拉特德拉特Kalaatderat
}
    
func (c *巴士拉BasraCounty) BMohammera莫哈梅拉() feud.Barony {
	return c.莫哈梅拉Mohammera
}
    
func (c *巴士拉BasraCounty) BQurna喀尔那() feud.Barony {
	return c.喀尔那Qurna
}
    
func (c *巴士拉BasraCounty) BSukelsheyuhk苏克舍赫克() feud.Barony {
	return c.苏克舍赫克Sukelsheyuhk
}
    
func (c *巴士拉BasraCounty) BUmmqasr乌姆盖萨尔() feud.Barony {
	return c.乌姆盖萨尔Ummqasr
}
    
var CBasra巴士拉 BasraCounty = &巴士拉BasraCounty{}

func init() {
	f := CBasra巴士拉.(*巴士拉BasraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "649",
		Title:     "basra",
		TitleName: "巴士拉",
		TitleCode: "c_basra",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉Arah = BArah阿拉
	f.阿拉Arah.SetParent(f)
	
	f.祖拜尔Azzubayr = BAzzubayr祖拜尔
	f.祖拜尔Azzubayr.SetParent(f)
	
	f.巴士拉Basra = BBasra巴士拉
	f.巴士拉Basra.SetParent(f)
	
	f.卡拉特德拉特Kalaatderat = BKalaatderat卡拉特德拉特
	f.卡拉特德拉特Kalaatderat.SetParent(f)
	
	f.莫哈梅拉Mohammera = BMohammera莫哈梅拉
	f.莫哈梅拉Mohammera.SetParent(f)
	
	f.喀尔那Qurna = BQurna喀尔那
	f.喀尔那Qurna.SetParent(f)
	
	f.苏克舍赫克Sukelsheyuhk = BSukelsheyuhk苏克舍赫克
	f.苏克舍赫克Sukelsheyuhk.SetParent(f)
	
	f.乌姆盖萨尔Ummqasr = BUmmqasr乌姆盖萨尔
	f.乌姆盖萨尔Ummqasr.SetParent(f)
	
}
