package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlaniaCounty interface {
    feud.County
    BArkhyz阿尔黑兹() 	feud.Barony
    BMagas马加斯() 	feud.Barony
    BTamarasheni塔马拉舍尼() 	feud.Barony
    BTkhabayerdy特哈巴_叶尔德() 	feud.Barony
    BTskhinval茨欣瓦利() 	feud.Barony
    BVovnushki沃夫努什基() 	feud.Barony
    BZadalesk扎达列斯克() 	feud.Barony
    BZaur扎乌尔() 	feud.Barony
}

type 阿兰尼亚AlaniaCounty struct {
	feud.BaseCounty
	阿尔黑兹Arkhyz 	feud.Barony
	马加斯Magas 	feud.Barony
	塔马拉舍尼Tamarasheni 	feud.Barony
	特哈巴_叶尔德Tkhabayerdy 	feud.Barony
	茨欣瓦利Tskhinval 	feud.Barony
	沃夫努什基Vovnushki 	feud.Barony
	扎达列斯克Zadalesk 	feud.Barony
	扎乌尔Zaur 	feud.Barony
}

func (c *阿兰尼亚AlaniaCounty) BArkhyz阿尔黑兹() feud.Barony {
	return c.阿尔黑兹Arkhyz
}
    
func (c *阿兰尼亚AlaniaCounty) BMagas马加斯() feud.Barony {
	return c.马加斯Magas
}
    
func (c *阿兰尼亚AlaniaCounty) BTamarasheni塔马拉舍尼() feud.Barony {
	return c.塔马拉舍尼Tamarasheni
}
    
func (c *阿兰尼亚AlaniaCounty) BTkhabayerdy特哈巴_叶尔德() feud.Barony {
	return c.特哈巴_叶尔德Tkhabayerdy
}
    
func (c *阿兰尼亚AlaniaCounty) BTskhinval茨欣瓦利() feud.Barony {
	return c.茨欣瓦利Tskhinval
}
    
func (c *阿兰尼亚AlaniaCounty) BVovnushki沃夫努什基() feud.Barony {
	return c.沃夫努什基Vovnushki
}
    
func (c *阿兰尼亚AlaniaCounty) BZadalesk扎达列斯克() feud.Barony {
	return c.扎达列斯克Zadalesk
}
    
func (c *阿兰尼亚AlaniaCounty) BZaur扎乌尔() feud.Barony {
	return c.扎乌尔Zaur
}
    
var CAlania阿兰尼亚 AlaniaCounty = &阿兰尼亚AlaniaCounty{}

func init() {
	f := CAlania阿兰尼亚.(*阿兰尼亚AlaniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "603",
		Title:     "alania",
		TitleName: "阿兰尼亚",
		TitleCode: "c_alania",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔黑兹Arkhyz = BArkhyz阿尔黑兹
	f.阿尔黑兹Arkhyz.SetParent(f)
	
	f.马加斯Magas = BMagas马加斯
	f.马加斯Magas.SetParent(f)
	
	f.塔马拉舍尼Tamarasheni = BTamarasheni塔马拉舍尼
	f.塔马拉舍尼Tamarasheni.SetParent(f)
	
	f.特哈巴_叶尔德Tkhabayerdy = BTkhabayerdy特哈巴_叶尔德
	f.特哈巴_叶尔德Tkhabayerdy.SetParent(f)
	
	f.茨欣瓦利Tskhinval = BTskhinval茨欣瓦利
	f.茨欣瓦利Tskhinval.SetParent(f)
	
	f.沃夫努什基Vovnushki = BVovnushki沃夫努什基
	f.沃夫努什基Vovnushki.SetParent(f)
	
	f.扎达列斯克Zadalesk = BZadalesk扎达列斯克
	f.扎达列斯克Zadalesk.SetParent(f)
	
	f.扎乌尔Zaur = BZaur扎乌尔
	f.扎乌尔Zaur.SetParent(f)
	
}
