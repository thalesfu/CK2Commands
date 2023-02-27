package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BjarmiaCounty interface {
    feud.County
    BIntsy因齐() 	feud.Barony
    BKamenka_bj卡缅卡() 	feud.Barony
    BKepino克皮诺() 	feud.Barony
    BKusnezowa库兹涅佐瓦() 	feud.Barony
    BLampozhnya兰波日尼亚() 	feud.Barony
    BOkladnikowa奥克拉德尼科瓦() 	feud.Barony
    BOkulovsky奥库洛夫斯基() 	feud.Barony
}

type 比亚尔米亚BjarmiaCounty struct {
	feud.BaseCounty
	因齐Intsy 	feud.Barony
	卡缅卡Kamenka_bj 	feud.Barony
	克皮诺Kepino 	feud.Barony
	库兹涅佐瓦Kusnezowa 	feud.Barony
	兰波日尼亚Lampozhnya 	feud.Barony
	奥克拉德尼科瓦Okladnikowa 	feud.Barony
	奥库洛夫斯基Okulovsky 	feud.Barony
}

func (c *比亚尔米亚BjarmiaCounty) BIntsy因齐() feud.Barony {
	return c.因齐Intsy
}
    
func (c *比亚尔米亚BjarmiaCounty) BKamenka_bj卡缅卡() feud.Barony {
	return c.卡缅卡Kamenka_bj
}
    
func (c *比亚尔米亚BjarmiaCounty) BKepino克皮诺() feud.Barony {
	return c.克皮诺Kepino
}
    
func (c *比亚尔米亚BjarmiaCounty) BKusnezowa库兹涅佐瓦() feud.Barony {
	return c.库兹涅佐瓦Kusnezowa
}
    
func (c *比亚尔米亚BjarmiaCounty) BLampozhnya兰波日尼亚() feud.Barony {
	return c.兰波日尼亚Lampozhnya
}
    
func (c *比亚尔米亚BjarmiaCounty) BOkladnikowa奥克拉德尼科瓦() feud.Barony {
	return c.奥克拉德尼科瓦Okladnikowa
}
    
func (c *比亚尔米亚BjarmiaCounty) BOkulovsky奥库洛夫斯基() feud.Barony {
	return c.奥库洛夫斯基Okulovsky
}
    
var CBjarmia比亚尔米亚 BjarmiaCounty = &比亚尔米亚BjarmiaCounty{}

func init() {
	f := CBjarmia比亚尔米亚.(*比亚尔米亚BjarmiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "396",
		Title:     "bjarmia",
		TitleName: "比亚尔米亚",
		TitleCode: "c_bjarmia",
		Baronies:  map[string]feud.Barony{},
	}

	f.因齐Intsy = BIntsy因齐
	f.因齐Intsy.SetParent(f)
	
	f.卡缅卡Kamenka_bj = BKamenka_bj卡缅卡
	f.卡缅卡Kamenka_bj.SetParent(f)
	
	f.克皮诺Kepino = BKepino克皮诺
	f.克皮诺Kepino.SetParent(f)
	
	f.库兹涅佐瓦Kusnezowa = BKusnezowa库兹涅佐瓦
	f.库兹涅佐瓦Kusnezowa.SetParent(f)
	
	f.兰波日尼亚Lampozhnya = BLampozhnya兰波日尼亚
	f.兰波日尼亚Lampozhnya.SetParent(f)
	
	f.奥克拉德尼科瓦Okladnikowa = BOkladnikowa奥克拉德尼科瓦
	f.奥克拉德尼科瓦Okladnikowa.SetParent(f)
	
	f.奥库洛夫斯基Okulovsky = BOkulovsky奥库洛夫斯基
	f.奥库洛夫斯基Okulovsky.SetParent(f)
	
}
