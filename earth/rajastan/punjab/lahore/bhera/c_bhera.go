package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BheraCounty interface {
    feud.County
    BBehak_maken贝哈克马肯() 	feud.Barony
    BBhera吠罗() 	feud.Barony
    BGirjakh吉迦卡() 	feud.Barony
    BGorawas瞿罗婆斯() 	feud.Barony
    BKatasraj泪眼() 	feud.Barony
    BPhalia费里亚() 	feud.Barony
    BRampur罗摩补罗() 	feud.Barony
}

type 吠罗BheraCounty struct {
	feud.BaseCounty
	贝哈克马肯Behak_maken 	feud.Barony
	吠罗Bhera 	feud.Barony
	吉迦卡Girjakh 	feud.Barony
	瞿罗婆斯Gorawas 	feud.Barony
	泪眼Katasraj 	feud.Barony
	费里亚Phalia 	feud.Barony
	罗摩补罗Rampur 	feud.Barony
}

func (c *吠罗BheraCounty) BBehak_maken贝哈克马肯() feud.Barony {
	return c.贝哈克马肯Behak_maken
}
    
func (c *吠罗BheraCounty) BBhera吠罗() feud.Barony {
	return c.吠罗Bhera
}
    
func (c *吠罗BheraCounty) BGirjakh吉迦卡() feud.Barony {
	return c.吉迦卡Girjakh
}
    
func (c *吠罗BheraCounty) BGorawas瞿罗婆斯() feud.Barony {
	return c.瞿罗婆斯Gorawas
}
    
func (c *吠罗BheraCounty) BKatasraj泪眼() feud.Barony {
	return c.泪眼Katasraj
}
    
func (c *吠罗BheraCounty) BPhalia费里亚() feud.Barony {
	return c.费里亚Phalia
}
    
func (c *吠罗BheraCounty) BRampur罗摩补罗() feud.Barony {
	return c.罗摩补罗Rampur
}
    
var CBhera吠罗 BheraCounty = &吠罗BheraCounty{}

func init() {
	f := CBhera吠罗.(*吠罗BheraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1192",
		Title:     "bhera",
		TitleName: "吠罗",
		TitleCode: "c_bhera",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝哈克马肯Behak_maken = BBehak_maken贝哈克马肯
	f.贝哈克马肯Behak_maken.SetParent(f)
	
	f.吠罗Bhera = BBhera吠罗
	f.吠罗Bhera.SetParent(f)
	
	f.吉迦卡Girjakh = BGirjakh吉迦卡
	f.吉迦卡Girjakh.SetParent(f)
	
	f.瞿罗婆斯Gorawas = BGorawas瞿罗婆斯
	f.瞿罗婆斯Gorawas.SetParent(f)
	
	f.泪眼Katasraj = BKatasraj泪眼
	f.泪眼Katasraj.SetParent(f)
	
	f.费里亚Phalia = BPhalia费里亚
	f.费里亚Phalia.SetParent(f)
	
	f.罗摩补罗Rampur = BRampur罗摩补罗
	f.罗摩补罗Rampur.SetParent(f)
	
}
