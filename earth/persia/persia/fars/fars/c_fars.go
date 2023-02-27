package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FarsCounty interface {
    feud.County
    BDarab达拉卜() 	feud.Barony
    BEstahbanat埃斯塔赫巴纳特() 	feud.Barony
    BGerash盖拉什() 	feud.Barony
    BJahrom贾赫罗姆() 	feud.Barony
    BKakhesasan卡赫斯撒恩() 	feud.Barony
    BKhafr哈夫尔() 	feud.Barony
    BLamerd拉默尔德() 	feud.Barony
    BPerozabad卑路扎巴德() 	feud.Barony
}

type 达什特斯坦FarsCounty struct {
	feud.BaseCounty
	达拉卜Darab 	feud.Barony
	埃斯塔赫巴纳特Estahbanat 	feud.Barony
	盖拉什Gerash 	feud.Barony
	贾赫罗姆Jahrom 	feud.Barony
	卡赫斯撒恩Kakhesasan 	feud.Barony
	哈夫尔Khafr 	feud.Barony
	拉默尔德Lamerd 	feud.Barony
	卑路扎巴德Perozabad 	feud.Barony
}

func (c *达什特斯坦FarsCounty) BDarab达拉卜() feud.Barony {
	return c.达拉卜Darab
}
    
func (c *达什特斯坦FarsCounty) BEstahbanat埃斯塔赫巴纳特() feud.Barony {
	return c.埃斯塔赫巴纳特Estahbanat
}
    
func (c *达什特斯坦FarsCounty) BGerash盖拉什() feud.Barony {
	return c.盖拉什Gerash
}
    
func (c *达什特斯坦FarsCounty) BJahrom贾赫罗姆() feud.Barony {
	return c.贾赫罗姆Jahrom
}
    
func (c *达什特斯坦FarsCounty) BKakhesasan卡赫斯撒恩() feud.Barony {
	return c.卡赫斯撒恩Kakhesasan
}
    
func (c *达什特斯坦FarsCounty) BKhafr哈夫尔() feud.Barony {
	return c.哈夫尔Khafr
}
    
func (c *达什特斯坦FarsCounty) BLamerd拉默尔德() feud.Barony {
	return c.拉默尔德Lamerd
}
    
func (c *达什特斯坦FarsCounty) BPerozabad卑路扎巴德() feud.Barony {
	return c.卑路扎巴德Perozabad
}
    
var CFars达什特斯坦 FarsCounty = &达什特斯坦FarsCounty{}

func init() {
	f := CFars达什特斯坦.(*达什特斯坦FarsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "643",
		Title:     "fars",
		TitleName: "达什特斯坦",
		TitleCode: "c_fars",
		Baronies:  map[string]feud.Barony{},
	}

	f.达拉卜Darab = BDarab达拉卜
	f.达拉卜Darab.SetParent(f)
	
	f.埃斯塔赫巴纳特Estahbanat = BEstahbanat埃斯塔赫巴纳特
	f.埃斯塔赫巴纳特Estahbanat.SetParent(f)
	
	f.盖拉什Gerash = BGerash盖拉什
	f.盖拉什Gerash.SetParent(f)
	
	f.贾赫罗姆Jahrom = BJahrom贾赫罗姆
	f.贾赫罗姆Jahrom.SetParent(f)
	
	f.卡赫斯撒恩Kakhesasan = BKakhesasan卡赫斯撒恩
	f.卡赫斯撒恩Kakhesasan.SetParent(f)
	
	f.哈夫尔Khafr = BKhafr哈夫尔
	f.哈夫尔Khafr.SetParent(f)
	
	f.拉默尔德Lamerd = BLamerd拉默尔德
	f.拉默尔德Lamerd.SetParent(f)
	
	f.卑路扎巴德Perozabad = BPerozabad卑路扎巴德
	f.卑路扎巴德Perozabad.SetParent(f)
	
}
