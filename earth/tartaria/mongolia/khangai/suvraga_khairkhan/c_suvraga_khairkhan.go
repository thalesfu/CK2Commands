package suvraga_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Suvraga_khairkhanCounty interface {
    feud.County
    BJargalant_suvraga扎尔嘎朗特() 	feud.Barony
    BKhotont赫吞特() 	feud.Barony
    BKhushuut_suvraga赫舍特() 	feud.Barony
    BMogod毛高德() 	feud.Barony
    BSuvraga_khairkhan苏布日嘎海尔汗() 	feud.Barony
    BTsenher臣赫尔() 	feud.Barony
    BZegstei哲格苏台() 	feud.Barony
}

type 苏布日嘎海尔汗Suvraga_khairkhanCounty struct {
	feud.BaseCounty
	扎尔嘎朗特Jargalant_suvraga 	feud.Barony
	赫吞特Khotont 	feud.Barony
	赫舍特Khushuut_suvraga 	feud.Barony
	毛高德Mogod 	feud.Barony
	苏布日嘎海尔汗Suvraga_khairkhan 	feud.Barony
	臣赫尔Tsenher 	feud.Barony
	哲格苏台Zegstei 	feud.Barony
}

func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BJargalant_suvraga扎尔嘎朗特() feud.Barony {
	return c.扎尔嘎朗特Jargalant_suvraga
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BKhotont赫吞特() feud.Barony {
	return c.赫吞特Khotont
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BKhushuut_suvraga赫舍特() feud.Barony {
	return c.赫舍特Khushuut_suvraga
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BMogod毛高德() feud.Barony {
	return c.毛高德Mogod
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BSuvraga_khairkhan苏布日嘎海尔汗() feud.Barony {
	return c.苏布日嘎海尔汗Suvraga_khairkhan
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BTsenher臣赫尔() feud.Barony {
	return c.臣赫尔Tsenher
}
    
func (c *苏布日嘎海尔汗Suvraga_khairkhanCounty) BZegstei哲格苏台() feud.Barony {
	return c.哲格苏台Zegstei
}
    
var CSuvraga_khairkhan苏布日嘎海尔汗 Suvraga_khairkhanCounty = &苏布日嘎海尔汗Suvraga_khairkhanCounty{}

func init() {
	f := CSuvraga_khairkhan苏布日嘎海尔汗.(*苏布日嘎海尔汗Suvraga_khairkhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1913",
		Title:     "suvraga_khairkhan",
		TitleName: "苏布日嘎海尔汗",
		TitleCode: "c_suvraga_khairkhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.扎尔嘎朗特Jargalant_suvraga = BJargalant_suvraga扎尔嘎朗特
	f.扎尔嘎朗特Jargalant_suvraga.SetParent(f)
	
	f.赫吞特Khotont = BKhotont赫吞特
	f.赫吞特Khotont.SetParent(f)
	
	f.赫舍特Khushuut_suvraga = BKhushuut_suvraga赫舍特
	f.赫舍特Khushuut_suvraga.SetParent(f)
	
	f.毛高德Mogod = BMogod毛高德
	f.毛高德Mogod.SetParent(f)
	
	f.苏布日嘎海尔汗Suvraga_khairkhan = BSuvraga_khairkhan苏布日嘎海尔汗
	f.苏布日嘎海尔汗Suvraga_khairkhan.SetParent(f)
	
	f.臣赫尔Tsenher = BTsenher臣赫尔
	f.臣赫尔Tsenher.SetParent(f)
	
	f.哲格苏台Zegstei = BZegstei哲格苏台
	f.哲格苏台Zegstei.SetParent(f)
	
}
