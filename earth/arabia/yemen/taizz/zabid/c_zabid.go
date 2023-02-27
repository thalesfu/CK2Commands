package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZabidCounty interface {
    feud.County
    BAl_mahjam马赫贾姆() 	feud.Barony
    BFasal法萨利() 	feud.Barony
    BGhalafiqa盖拉菲加() 	feud.Barony
    BHays海斯() 	feud.Barony
    BMawr毛尔() 	feud.Barony
    BMukha穆哈() 	feud.Barony
    BZabid宰比德() 	feud.Barony
}

type 宰比德ZabidCounty struct {
	feud.BaseCounty
	马赫贾姆Al_mahjam 	feud.Barony
	法萨利Fasal 	feud.Barony
	盖拉菲加Ghalafiqa 	feud.Barony
	海斯Hays 	feud.Barony
	毛尔Mawr 	feud.Barony
	穆哈Mukha 	feud.Barony
	宰比德Zabid 	feud.Barony
}

func (c *宰比德ZabidCounty) BAl_mahjam马赫贾姆() feud.Barony {
	return c.马赫贾姆Al_mahjam
}
    
func (c *宰比德ZabidCounty) BFasal法萨利() feud.Barony {
	return c.法萨利Fasal
}
    
func (c *宰比德ZabidCounty) BGhalafiqa盖拉菲加() feud.Barony {
	return c.盖拉菲加Ghalafiqa
}
    
func (c *宰比德ZabidCounty) BHays海斯() feud.Barony {
	return c.海斯Hays
}
    
func (c *宰比德ZabidCounty) BMawr毛尔() feud.Barony {
	return c.毛尔Mawr
}
    
func (c *宰比德ZabidCounty) BMukha穆哈() feud.Barony {
	return c.穆哈Mukha
}
    
func (c *宰比德ZabidCounty) BZabid宰比德() feud.Barony {
	return c.宰比德Zabid
}
    
var CZabid宰比德 ZabidCounty = &宰比德ZabidCounty{}

func init() {
	f := CZabid宰比德.(*宰比德ZabidCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1532",
		Title:     "zabid",
		TitleName: "宰比德",
		TitleCode: "c_zabid",
		Baronies:  map[string]feud.Barony{},
	}

	f.马赫贾姆Al_mahjam = BAl_mahjam马赫贾姆
	f.马赫贾姆Al_mahjam.SetParent(f)
	
	f.法萨利Fasal = BFasal法萨利
	f.法萨利Fasal.SetParent(f)
	
	f.盖拉菲加Ghalafiqa = BGhalafiqa盖拉菲加
	f.盖拉菲加Ghalafiqa.SetParent(f)
	
	f.海斯Hays = BHays海斯
	f.海斯Hays.SetParent(f)
	
	f.毛尔Mawr = BMawr毛尔
	f.毛尔Mawr.SetParent(f)
	
	f.穆哈Mukha = BMukha穆哈
	f.穆哈Mukha.SetParent(f)
	
	f.宰比德Zabid = BZabid宰比德
	f.宰比德Zabid.SetParent(f)
	
}
