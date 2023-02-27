package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MordvaCounty interface {
    feud.County
    BArdatovo阿尔达托沃() 	feud.Barony
    BInsar因萨尔() 	feud.Barony
    BKrasnoslobodsk克拉斯诺斯洛博茨克() 	feud.Barony
    BPenza奔萨() 	feud.Barony
    BSaransk萨兰斯克() 	feud.Barony
    BTemnikow捷姆尼科夫() 	feud.Barony
    BYalga亚尔加() 	feud.Barony
    BYavas亚瓦斯() 	feud.Barony
}

type 莫尔多瓦MordvaCounty struct {
	feud.BaseCounty
	阿尔达托沃Ardatovo 	feud.Barony
	因萨尔Insar 	feud.Barony
	克拉斯诺斯洛博茨克Krasnoslobodsk 	feud.Barony
	奔萨Penza 	feud.Barony
	萨兰斯克Saransk 	feud.Barony
	捷姆尼科夫Temnikow 	feud.Barony
	亚尔加Yalga 	feud.Barony
	亚瓦斯Yavas 	feud.Barony
}

func (c *莫尔多瓦MordvaCounty) BArdatovo阿尔达托沃() feud.Barony {
	return c.阿尔达托沃Ardatovo
}
    
func (c *莫尔多瓦MordvaCounty) BInsar因萨尔() feud.Barony {
	return c.因萨尔Insar
}
    
func (c *莫尔多瓦MordvaCounty) BKrasnoslobodsk克拉斯诺斯洛博茨克() feud.Barony {
	return c.克拉斯诺斯洛博茨克Krasnoslobodsk
}
    
func (c *莫尔多瓦MordvaCounty) BPenza奔萨() feud.Barony {
	return c.奔萨Penza
}
    
func (c *莫尔多瓦MordvaCounty) BSaransk萨兰斯克() feud.Barony {
	return c.萨兰斯克Saransk
}
    
func (c *莫尔多瓦MordvaCounty) BTemnikow捷姆尼科夫() feud.Barony {
	return c.捷姆尼科夫Temnikow
}
    
func (c *莫尔多瓦MordvaCounty) BYalga亚尔加() feud.Barony {
	return c.亚尔加Yalga
}
    
func (c *莫尔多瓦MordvaCounty) BYavas亚瓦斯() feud.Barony {
	return c.亚瓦斯Yavas
}
    
var CMordva莫尔多瓦 MordvaCounty = &莫尔多瓦MordvaCounty{}

func init() {
	f := CMordva莫尔多瓦.(*莫尔多瓦MordvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "579",
		Title:     "mordva",
		TitleName: "莫尔多瓦",
		TitleCode: "c_mordva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔达托沃Ardatovo = BArdatovo阿尔达托沃
	f.阿尔达托沃Ardatovo.SetParent(f)
	
	f.因萨尔Insar = BInsar因萨尔
	f.因萨尔Insar.SetParent(f)
	
	f.克拉斯诺斯洛博茨克Krasnoslobodsk = BKrasnoslobodsk克拉斯诺斯洛博茨克
	f.克拉斯诺斯洛博茨克Krasnoslobodsk.SetParent(f)
	
	f.奔萨Penza = BPenza奔萨
	f.奔萨Penza.SetParent(f)
	
	f.萨兰斯克Saransk = BSaransk萨兰斯克
	f.萨兰斯克Saransk.SetParent(f)
	
	f.捷姆尼科夫Temnikow = BTemnikow捷姆尼科夫
	f.捷姆尼科夫Temnikow.SetParent(f)
	
	f.亚尔加Yalga = BYalga亚尔加
	f.亚尔加Yalga.SetParent(f)
	
	f.亚瓦斯Yavas = BYavas亚瓦斯
	f.亚瓦斯Yavas.SetParent(f)
	
}
