package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AvhazCounty interface {
    feud.County
    BDezful迪兹富勒() 	feud.Barony
    BFarsan法尔桑() 	feud.Barony
    BHelen赫伦() 	feud.Barony
    BIdhaj伊泽() 	feud.Barony
    BKoohrang库兰格() 	feud.Barony
    BMasjedsoleyman马斯吉德苏莱曼() 	feud.Barony
    BShahrekord沙赫尔库尔德() 	feud.Barony
    BShushtar舒什塔尔() 	feud.Barony
}

type 阿瓦士AvhazCounty struct {
	feud.BaseCounty
	迪兹富勒Dezful 	feud.Barony
	法尔桑Farsan 	feud.Barony
	赫伦Helen 	feud.Barony
	伊泽Idhaj 	feud.Barony
	库兰格Koohrang 	feud.Barony
	马斯吉德苏莱曼Masjedsoleyman 	feud.Barony
	沙赫尔库尔德Shahrekord 	feud.Barony
	舒什塔尔Shushtar 	feud.Barony
}

func (c *阿瓦士AvhazCounty) BDezful迪兹富勒() feud.Barony {
	return c.迪兹富勒Dezful
}
    
func (c *阿瓦士AvhazCounty) BFarsan法尔桑() feud.Barony {
	return c.法尔桑Farsan
}
    
func (c *阿瓦士AvhazCounty) BHelen赫伦() feud.Barony {
	return c.赫伦Helen
}
    
func (c *阿瓦士AvhazCounty) BIdhaj伊泽() feud.Barony {
	return c.伊泽Idhaj
}
    
func (c *阿瓦士AvhazCounty) BKoohrang库兰格() feud.Barony {
	return c.库兰格Koohrang
}
    
func (c *阿瓦士AvhazCounty) BMasjedsoleyman马斯吉德苏莱曼() feud.Barony {
	return c.马斯吉德苏莱曼Masjedsoleyman
}
    
func (c *阿瓦士AvhazCounty) BShahrekord沙赫尔库尔德() feud.Barony {
	return c.沙赫尔库尔德Shahrekord
}
    
func (c *阿瓦士AvhazCounty) BShushtar舒什塔尔() feud.Barony {
	return c.舒什塔尔Shushtar
}
    
var CAvhaz阿瓦士 AvhazCounty = &阿瓦士AvhazCounty{}

func init() {
	f := CAvhaz阿瓦士.(*阿瓦士AvhazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "647",
		Title:     "avhaz",
		TitleName: "阿瓦士",
		TitleCode: "c_avhaz",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪兹富勒Dezful = BDezful迪兹富勒
	f.迪兹富勒Dezful.SetParent(f)
	
	f.法尔桑Farsan = BFarsan法尔桑
	f.法尔桑Farsan.SetParent(f)
	
	f.赫伦Helen = BHelen赫伦
	f.赫伦Helen.SetParent(f)
	
	f.伊泽Idhaj = BIdhaj伊泽
	f.伊泽Idhaj.SetParent(f)
	
	f.库兰格Koohrang = BKoohrang库兰格
	f.库兰格Koohrang.SetParent(f)
	
	f.马斯吉德苏莱曼Masjedsoleyman = BMasjedsoleyman马斯吉德苏莱曼
	f.马斯吉德苏莱曼Masjedsoleyman.SetParent(f)
	
	f.沙赫尔库尔德Shahrekord = BShahrekord沙赫尔库尔德
	f.沙赫尔库尔德Shahrekord.SetParent(f)
	
	f.舒什塔尔Shushtar = BShushtar舒什塔尔
	f.舒什塔尔Shushtar.SetParent(f)
	
}
