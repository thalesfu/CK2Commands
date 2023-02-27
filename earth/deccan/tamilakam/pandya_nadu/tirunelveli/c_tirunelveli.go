package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TirunelveliCounty interface {
    feud.County
    BKalugumalai伽卢古摩罗() 	feud.Barony
    BKayal伽一() 	feud.Barony
    BKayattar迦耶多() 	feud.Barony
    BKorkai拘罗盖() 	feud.Barony
    BSantera桑谛罗() 	feud.Barony
    BTirunelveli底卢内勒吠梨() 	feud.Barony
    BTiruvalisvaram帝卢伐利湿伐蓝() 	feud.Barony
}

type 底卢内勒吠梨TirunelveliCounty struct {
	feud.BaseCounty
	伽卢古摩罗Kalugumalai 	feud.Barony
	伽一Kayal 	feud.Barony
	迦耶多Kayattar 	feud.Barony
	拘罗盖Korkai 	feud.Barony
	桑谛罗Santera 	feud.Barony
	底卢内勒吠梨Tirunelveli 	feud.Barony
	帝卢伐利湿伐蓝Tiruvalisvaram 	feud.Barony
}

func (c *底卢内勒吠梨TirunelveliCounty) BKalugumalai伽卢古摩罗() feud.Barony {
	return c.伽卢古摩罗Kalugumalai
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BKayal伽一() feud.Barony {
	return c.伽一Kayal
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BKayattar迦耶多() feud.Barony {
	return c.迦耶多Kayattar
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BKorkai拘罗盖() feud.Barony {
	return c.拘罗盖Korkai
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BSantera桑谛罗() feud.Barony {
	return c.桑谛罗Santera
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BTirunelveli底卢内勒吠梨() feud.Barony {
	return c.底卢内勒吠梨Tirunelveli
}
    
func (c *底卢内勒吠梨TirunelveliCounty) BTiruvalisvaram帝卢伐利湿伐蓝() feud.Barony {
	return c.帝卢伐利湿伐蓝Tiruvalisvaram
}
    
var CTirunelveli底卢内勒吠梨 TirunelveliCounty = &底卢内勒吠梨TirunelveliCounty{}

func init() {
	f := CTirunelveli底卢内勒吠梨.(*底卢内勒吠梨TirunelveliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1214",
		Title:     "tirunelveli",
		TitleName: "底卢内勒吠梨",
		TitleCode: "c_tirunelveli",
		Baronies:  map[string]feud.Barony{},
	}

	f.伽卢古摩罗Kalugumalai = BKalugumalai伽卢古摩罗
	f.伽卢古摩罗Kalugumalai.SetParent(f)
	
	f.伽一Kayal = BKayal伽一
	f.伽一Kayal.SetParent(f)
	
	f.迦耶多Kayattar = BKayattar迦耶多
	f.迦耶多Kayattar.SetParent(f)
	
	f.拘罗盖Korkai = BKorkai拘罗盖
	f.拘罗盖Korkai.SetParent(f)
	
	f.桑谛罗Santera = BSantera桑谛罗
	f.桑谛罗Santera.SetParent(f)
	
	f.底卢内勒吠梨Tirunelveli = BTirunelveli底卢内勒吠梨
	f.底卢内勒吠梨Tirunelveli.SetParent(f)
	
	f.帝卢伐利湿伐蓝Tiruvalisvaram = BTiruvalisvaram帝卢伐利湿伐蓝
	f.帝卢伐利湿伐蓝Tiruvalisvaram.SetParent(f)
	
}
