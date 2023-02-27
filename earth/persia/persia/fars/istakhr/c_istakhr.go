package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IstakhrCounty interface {
    feud.County
    BAbadha阿巴扎() 	feud.Barony
    BAbarquh阿巴尔古() 	feud.Barony
    BIqlid伊格利德() 	feud.Barony
    BIstakhr伊斯塔赫尔() 	feud.Barony
    BKamin卡明() 	feud.Barony
    BMain马因() 	feud.Barony
    BSarmaq萨尔马格() 	feud.Barony
}

type 伊斯塔赫尔IstakhrCounty struct {
	feud.BaseCounty
	阿巴扎Abadha 	feud.Barony
	阿巴尔古Abarquh 	feud.Barony
	伊格利德Iqlid 	feud.Barony
	伊斯塔赫尔Istakhr 	feud.Barony
	卡明Kamin 	feud.Barony
	马因Main 	feud.Barony
	萨尔马格Sarmaq 	feud.Barony
}

func (c *伊斯塔赫尔IstakhrCounty) BAbadha阿巴扎() feud.Barony {
	return c.阿巴扎Abadha
}
    
func (c *伊斯塔赫尔IstakhrCounty) BAbarquh阿巴尔古() feud.Barony {
	return c.阿巴尔古Abarquh
}
    
func (c *伊斯塔赫尔IstakhrCounty) BIqlid伊格利德() feud.Barony {
	return c.伊格利德Iqlid
}
    
func (c *伊斯塔赫尔IstakhrCounty) BIstakhr伊斯塔赫尔() feud.Barony {
	return c.伊斯塔赫尔Istakhr
}
    
func (c *伊斯塔赫尔IstakhrCounty) BKamin卡明() feud.Barony {
	return c.卡明Kamin
}
    
func (c *伊斯塔赫尔IstakhrCounty) BMain马因() feud.Barony {
	return c.马因Main
}
    
func (c *伊斯塔赫尔IstakhrCounty) BSarmaq萨尔马格() feud.Barony {
	return c.萨尔马格Sarmaq
}
    
var CIstakhr伊斯塔赫尔 IstakhrCounty = &伊斯塔赫尔IstakhrCounty{}

func init() {
	f := CIstakhr伊斯塔赫尔.(*伊斯塔赫尔IstakhrCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "691",
		Title:     "istakhr",
		TitleName: "伊斯塔赫尔",
		TitleCode: "c_istakhr",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴扎Abadha = BAbadha阿巴扎
	f.阿巴扎Abadha.SetParent(f)
	
	f.阿巴尔古Abarquh = BAbarquh阿巴尔古
	f.阿巴尔古Abarquh.SetParent(f)
	
	f.伊格利德Iqlid = BIqlid伊格利德
	f.伊格利德Iqlid.SetParent(f)
	
	f.伊斯塔赫尔Istakhr = BIstakhr伊斯塔赫尔
	f.伊斯塔赫尔Istakhr.SetParent(f)
	
	f.卡明Kamin = BKamin卡明
	f.卡明Kamin.SetParent(f)
	
	f.马因Main = BMain马因
	f.马因Main.SetParent(f)
	
	f.萨尔马格Sarmaq = BSarmaq萨尔马格
	f.萨尔马格Sarmaq.SetParent(f)
	
}
