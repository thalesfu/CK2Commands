package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NabadwipaCounty interface {
    feud.County
    BBidhyanandikathi毗陀耶难提迦底() 	feud.Barony
    BGodrumdwip瞿度摩洲() 	feud.Barony
    BKrishnagar奎师那揭罗() 	feud.Barony
    BMadhyadwip摩地耶洲() 	feud.Barony
    BNabadwipa那婆提鞞波() 	feud.Barony
    BPancharatna般遮剌那() 	feud.Barony
    BRitudwip利堵洲() 	feud.Barony
}

type 那婆提鞞波NabadwipaCounty struct {
	feud.BaseCounty
	毗陀耶难提迦底Bidhyanandikathi 	feud.Barony
	瞿度摩洲Godrumdwip 	feud.Barony
	奎师那揭罗Krishnagar 	feud.Barony
	摩地耶洲Madhyadwip 	feud.Barony
	那婆提鞞波Nabadwipa 	feud.Barony
	般遮剌那Pancharatna 	feud.Barony
	利堵洲Ritudwip 	feud.Barony
}

func (c *那婆提鞞波NabadwipaCounty) BBidhyanandikathi毗陀耶难提迦底() feud.Barony {
	return c.毗陀耶难提迦底Bidhyanandikathi
}
    
func (c *那婆提鞞波NabadwipaCounty) BGodrumdwip瞿度摩洲() feud.Barony {
	return c.瞿度摩洲Godrumdwip
}
    
func (c *那婆提鞞波NabadwipaCounty) BKrishnagar奎师那揭罗() feud.Barony {
	return c.奎师那揭罗Krishnagar
}
    
func (c *那婆提鞞波NabadwipaCounty) BMadhyadwip摩地耶洲() feud.Barony {
	return c.摩地耶洲Madhyadwip
}
    
func (c *那婆提鞞波NabadwipaCounty) BNabadwipa那婆提鞞波() feud.Barony {
	return c.那婆提鞞波Nabadwipa
}
    
func (c *那婆提鞞波NabadwipaCounty) BPancharatna般遮剌那() feud.Barony {
	return c.般遮剌那Pancharatna
}
    
func (c *那婆提鞞波NabadwipaCounty) BRitudwip利堵洲() feud.Barony {
	return c.利堵洲Ritudwip
}
    
var CNabadwipa那婆提鞞波 NabadwipaCounty = &那婆提鞞波NabadwipaCounty{}

func init() {
	f := CNabadwipa那婆提鞞波.(*那婆提鞞波NabadwipaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1323",
		Title:     "nabadwipa",
		TitleName: "那婆提鞞波",
		TitleCode: "c_nabadwipa",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗陀耶难提迦底Bidhyanandikathi = BBidhyanandikathi毗陀耶难提迦底
	f.毗陀耶难提迦底Bidhyanandikathi.SetParent(f)
	
	f.瞿度摩洲Godrumdwip = BGodrumdwip瞿度摩洲
	f.瞿度摩洲Godrumdwip.SetParent(f)
	
	f.奎师那揭罗Krishnagar = BKrishnagar奎师那揭罗
	f.奎师那揭罗Krishnagar.SetParent(f)
	
	f.摩地耶洲Madhyadwip = BMadhyadwip摩地耶洲
	f.摩地耶洲Madhyadwip.SetParent(f)
	
	f.那婆提鞞波Nabadwipa = BNabadwipa那婆提鞞波
	f.那婆提鞞波Nabadwipa.SetParent(f)
	
	f.般遮剌那Pancharatna = BPancharatna般遮剌那
	f.般遮剌那Pancharatna.SetParent(f)
	
	f.利堵洲Ritudwip = BRitudwip利堵洲
	f.利堵洲Ritudwip.SetParent(f)
	
}
