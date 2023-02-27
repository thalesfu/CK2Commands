package farrah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FarrahCounty interface {
    feud.County
    BAnardara阿那尔达雷() 	feud.Barony
    BBakwa巴克瓦() 	feud.Barony
    BBaladuluk巴拉达鲁克() 	feud.Barony
    BFarrah法拉() 	feud.Barony
    BJuwayn朱韦恩() 	feud.Barony
    BKhakisafed哈基萨法德() 	feud.Barony
    BQalaikah奇拉卡赫() 	feud.Barony
    BShibkoh司阔霍() 	feud.Barony
}

type 法拉FarrahCounty struct {
	feud.BaseCounty
	阿那尔达雷Anardara 	feud.Barony
	巴克瓦Bakwa 	feud.Barony
	巴拉达鲁克Baladuluk 	feud.Barony
	法拉Farrah 	feud.Barony
	朱韦恩Juwayn 	feud.Barony
	哈基萨法德Khakisafed 	feud.Barony
	奇拉卡赫Qalaikah 	feud.Barony
	司阔霍Shibkoh 	feud.Barony
}

func (c *法拉FarrahCounty) BAnardara阿那尔达雷() feud.Barony {
	return c.阿那尔达雷Anardara
}
    
func (c *法拉FarrahCounty) BBakwa巴克瓦() feud.Barony {
	return c.巴克瓦Bakwa
}
    
func (c *法拉FarrahCounty) BBaladuluk巴拉达鲁克() feud.Barony {
	return c.巴拉达鲁克Baladuluk
}
    
func (c *法拉FarrahCounty) BFarrah法拉() feud.Barony {
	return c.法拉Farrah
}
    
func (c *法拉FarrahCounty) BJuwayn朱韦恩() feud.Barony {
	return c.朱韦恩Juwayn
}
    
func (c *法拉FarrahCounty) BKhakisafed哈基萨法德() feud.Barony {
	return c.哈基萨法德Khakisafed
}
    
func (c *法拉FarrahCounty) BQalaikah奇拉卡赫() feud.Barony {
	return c.奇拉卡赫Qalaikah
}
    
func (c *法拉FarrahCounty) BShibkoh司阔霍() feud.Barony {
	return c.司阔霍Shibkoh
}
    
var CFarrah法拉 FarrahCounty = &法拉FarrahCounty{}

func init() {
	f := CFarrah法拉.(*法拉FarrahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "299",
		Title:     "farrah",
		TitleName: "法拉",
		TitleCode: "c_farrah",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿那尔达雷Anardara = BAnardara阿那尔达雷
	f.阿那尔达雷Anardara.SetParent(f)
	
	f.巴克瓦Bakwa = BBakwa巴克瓦
	f.巴克瓦Bakwa.SetParent(f)
	
	f.巴拉达鲁克Baladuluk = BBaladuluk巴拉达鲁克
	f.巴拉达鲁克Baladuluk.SetParent(f)
	
	f.法拉Farrah = BFarrah法拉
	f.法拉Farrah.SetParent(f)
	
	f.朱韦恩Juwayn = BJuwayn朱韦恩
	f.朱韦恩Juwayn.SetParent(f)
	
	f.哈基萨法德Khakisafed = BKhakisafed哈基萨法德
	f.哈基萨法德Khakisafed.SetParent(f)
	
	f.奇拉卡赫Qalaikah = BQalaikah奇拉卡赫
	f.奇拉卡赫Qalaikah.SetParent(f)
	
	f.司阔霍Shibkoh = BShibkoh司阔霍
	f.司阔霍Shibkoh.SetParent(f)
	
}
