package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KirkukCounty interface {
    feud.County
    BChuartan初阿坦() 	feud.Barony
    BDaquq达古格() 	feud.Barony
    BDukan杜坎() 	feud.Barony
    BHalabja哈莱卜杰() 	feud.Barony
    BKifri基夫里() 	feud.Barony
    BKirkuk基尔库克() 	feud.Barony
    BMakhmur迈赫穆尔() 	feud.Barony
    BRanya拉尼耶() 	feud.Barony
}

type 基尔库克KirkukCounty struct {
	feud.BaseCounty
	初阿坦Chuartan 	feud.Barony
	达古格Daquq 	feud.Barony
	杜坎Dukan 	feud.Barony
	哈莱卜杰Halabja 	feud.Barony
	基夫里Kifri 	feud.Barony
	基尔库克Kirkuk 	feud.Barony
	迈赫穆尔Makhmur 	feud.Barony
	拉尼耶Ranya 	feud.Barony
}

func (c *基尔库克KirkukCounty) BChuartan初阿坦() feud.Barony {
	return c.初阿坦Chuartan
}
    
func (c *基尔库克KirkukCounty) BDaquq达古格() feud.Barony {
	return c.达古格Daquq
}
    
func (c *基尔库克KirkukCounty) BDukan杜坎() feud.Barony {
	return c.杜坎Dukan
}
    
func (c *基尔库克KirkukCounty) BHalabja哈莱卜杰() feud.Barony {
	return c.哈莱卜杰Halabja
}
    
func (c *基尔库克KirkukCounty) BKifri基夫里() feud.Barony {
	return c.基夫里Kifri
}
    
func (c *基尔库克KirkukCounty) BKirkuk基尔库克() feud.Barony {
	return c.基尔库克Kirkuk
}
    
func (c *基尔库克KirkukCounty) BMakhmur迈赫穆尔() feud.Barony {
	return c.迈赫穆尔Makhmur
}
    
func (c *基尔库克KirkukCounty) BRanya拉尼耶() feud.Barony {
	return c.拉尼耶Ranya
}
    
var CKirkuk基尔库克 KirkukCounty = &基尔库克KirkukCounty{}

func init() {
	f := CKirkuk基尔库克.(*基尔库克KirkukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "687",
		Title:     "kirkuk",
		TitleName: "基尔库克",
		TitleCode: "c_kirkuk",
		Baronies:  map[string]feud.Barony{},
	}

	f.初阿坦Chuartan = BChuartan初阿坦
	f.初阿坦Chuartan.SetParent(f)
	
	f.达古格Daquq = BDaquq达古格
	f.达古格Daquq.SetParent(f)
	
	f.杜坎Dukan = BDukan杜坎
	f.杜坎Dukan.SetParent(f)
	
	f.哈莱卜杰Halabja = BHalabja哈莱卜杰
	f.哈莱卜杰Halabja.SetParent(f)
	
	f.基夫里Kifri = BKifri基夫里
	f.基夫里Kifri.SetParent(f)
	
	f.基尔库克Kirkuk = BKirkuk基尔库克
	f.基尔库克Kirkuk.SetParent(f)
	
	f.迈赫穆尔Makhmur = BMakhmur迈赫穆尔
	f.迈赫穆尔Makhmur.SetParent(f)
	
	f.拉尼耶Ranya = BRanya拉尼耶
	f.拉尼耶Ranya.SetParent(f)
	
}
