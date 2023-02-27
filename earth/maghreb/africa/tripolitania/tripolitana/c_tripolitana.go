package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TripolitanaCounty interface {
    feud.County
    BAljamil阿贾米拉() 	feud.Barony
    BBaniwaled拜尼沃利德() 	feud.Barony
    BFunduq福杜格() 	feud.Barony
    BGherlan格赫尔兰() 	feud.Barony
    BLibtripoli的黎波里() 	feud.Barony
    BMazuzah马祖扎哈() 	feud.Barony
    BSabratah塞卜拉泰() 	feud.Barony
}

type 的黎波里塔尼亚TripolitanaCounty struct {
	feud.BaseCounty
	阿贾米拉Aljamil 	feud.Barony
	拜尼沃利德Baniwaled 	feud.Barony
	福杜格Funduq 	feud.Barony
	格赫尔兰Gherlan 	feud.Barony
	的黎波里Libtripoli 	feud.Barony
	马祖扎哈Mazuzah 	feud.Barony
	塞卜拉泰Sabratah 	feud.Barony
}

func (c *的黎波里塔尼亚TripolitanaCounty) BAljamil阿贾米拉() feud.Barony {
	return c.阿贾米拉Aljamil
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BBaniwaled拜尼沃利德() feud.Barony {
	return c.拜尼沃利德Baniwaled
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BFunduq福杜格() feud.Barony {
	return c.福杜格Funduq
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BGherlan格赫尔兰() feud.Barony {
	return c.格赫尔兰Gherlan
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BLibtripoli的黎波里() feud.Barony {
	return c.的黎波里Libtripoli
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BMazuzah马祖扎哈() feud.Barony {
	return c.马祖扎哈Mazuzah
}
    
func (c *的黎波里塔尼亚TripolitanaCounty) BSabratah塞卜拉泰() feud.Barony {
	return c.塞卜拉泰Sabratah
}
    
var CTripolitana的黎波里塔尼亚 TripolitanaCounty = &的黎波里塔尼亚TripolitanaCounty{}

func init() {
	f := CTripolitana的黎波里塔尼亚.(*的黎波里塔尼亚TripolitanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "811",
		Title:     "tripolitana",
		TitleName: "的黎波里塔尼亚",
		TitleCode: "c_tripolitana",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿贾米拉Aljamil = BAljamil阿贾米拉
	f.阿贾米拉Aljamil.SetParent(f)
	
	f.拜尼沃利德Baniwaled = BBaniwaled拜尼沃利德
	f.拜尼沃利德Baniwaled.SetParent(f)
	
	f.福杜格Funduq = BFunduq福杜格
	f.福杜格Funduq.SetParent(f)
	
	f.格赫尔兰Gherlan = BGherlan格赫尔兰
	f.格赫尔兰Gherlan.SetParent(f)
	
	f.的黎波里Libtripoli = BLibtripoli的黎波里
	f.的黎波里Libtripoli.SetParent(f)
	
	f.马祖扎哈Mazuzah = BMazuzah马祖扎哈
	f.马祖扎哈Mazuzah.SetParent(f)
	
	f.塞卜拉泰Sabratah = BSabratah塞卜拉泰
	f.塞卜拉泰Sabratah.SetParent(f)
	
}
