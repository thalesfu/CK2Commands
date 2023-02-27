package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KathmanduCounty interface {
    feud.County
    BBhaktapur薄多补罗() 	feud.Barony
    BKathmandu加德满都() 	feud.Barony
    BKirtipur诘帝补罗() 	feud.Barony
    BLalitpur罗梨多补罗() 	feud.Barony
    BNagarkot那揭罗拘吒() 	feud.Barony
    BPashupatinath波输波底那他() 	feud.Barony
    BPatan_yala帕坦亚拉() 	feud.Barony
}

type 加德满都KathmanduCounty struct {
	feud.BaseCounty
	薄多补罗Bhaktapur 	feud.Barony
	加德满都Kathmandu 	feud.Barony
	诘帝补罗Kirtipur 	feud.Barony
	罗梨多补罗Lalitpur 	feud.Barony
	那揭罗拘吒Nagarkot 	feud.Barony
	波输波底那他Pashupatinath 	feud.Barony
	帕坦亚拉Patan_yala 	feud.Barony
}

func (c *加德满都KathmanduCounty) BBhaktapur薄多补罗() feud.Barony {
	return c.薄多补罗Bhaktapur
}
    
func (c *加德满都KathmanduCounty) BKathmandu加德满都() feud.Barony {
	return c.加德满都Kathmandu
}
    
func (c *加德满都KathmanduCounty) BKirtipur诘帝补罗() feud.Barony {
	return c.诘帝补罗Kirtipur
}
    
func (c *加德满都KathmanduCounty) BLalitpur罗梨多补罗() feud.Barony {
	return c.罗梨多补罗Lalitpur
}
    
func (c *加德满都KathmanduCounty) BNagarkot那揭罗拘吒() feud.Barony {
	return c.那揭罗拘吒Nagarkot
}
    
func (c *加德满都KathmanduCounty) BPashupatinath波输波底那他() feud.Barony {
	return c.波输波底那他Pashupatinath
}
    
func (c *加德满都KathmanduCounty) BPatan_yala帕坦亚拉() feud.Barony {
	return c.帕坦亚拉Patan_yala
}
    
var CKathmandu加德满都 KathmanduCounty = &加德满都KathmanduCounty{}

func init() {
	f := CKathmandu加德满都.(*加德满都KathmanduCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1478",
		Title:     "kathmandu",
		TitleName: "加德满都",
		TitleCode: "c_kathmandu",
		Baronies:  map[string]feud.Barony{},
	}

	f.薄多补罗Bhaktapur = BBhaktapur薄多补罗
	f.薄多补罗Bhaktapur.SetParent(f)
	
	f.加德满都Kathmandu = BKathmandu加德满都
	f.加德满都Kathmandu.SetParent(f)
	
	f.诘帝补罗Kirtipur = BKirtipur诘帝补罗
	f.诘帝补罗Kirtipur.SetParent(f)
	
	f.罗梨多补罗Lalitpur = BLalitpur罗梨多补罗
	f.罗梨多补罗Lalitpur.SetParent(f)
	
	f.那揭罗拘吒Nagarkot = BNagarkot那揭罗拘吒
	f.那揭罗拘吒Nagarkot.SetParent(f)
	
	f.波输波底那他Pashupatinath = BPashupatinath波输波底那他
	f.波输波底那他Pashupatinath.SetParent(f)
	
	f.帕坦亚拉Patan_yala = BPatan_yala帕坦亚拉
	f.帕坦亚拉Patan_yala.SetParent(f)
	
}
