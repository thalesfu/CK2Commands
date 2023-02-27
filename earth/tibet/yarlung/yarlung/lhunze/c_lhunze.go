package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhunzeCounty interface {
    feud.County
    BCona_b错那() 	feud.Barony
    BLhunze隆子() 	feud.Barony
    BNariyong拿日雍() 	feud.Barony
    BRitang日当() 	feud.Barony
    BSangngagqoiling三安曲林() 	feud.Barony
    BYumai玉麦() 	feud.Barony
    BZhunba准巴() 	feud.Barony
}

type 隆子LhunzeCounty struct {
	feud.BaseCounty
	错那Cona_b 	feud.Barony
	隆子Lhunze 	feud.Barony
	拿日雍Nariyong 	feud.Barony
	日当Ritang 	feud.Barony
	三安曲林Sangngagqoiling 	feud.Barony
	玉麦Yumai 	feud.Barony
	准巴Zhunba 	feud.Barony
}

func (c *隆子LhunzeCounty) BCona_b错那() feud.Barony {
	return c.错那Cona_b
}
    
func (c *隆子LhunzeCounty) BLhunze隆子() feud.Barony {
	return c.隆子Lhunze
}
    
func (c *隆子LhunzeCounty) BNariyong拿日雍() feud.Barony {
	return c.拿日雍Nariyong
}
    
func (c *隆子LhunzeCounty) BRitang日当() feud.Barony {
	return c.日当Ritang
}
    
func (c *隆子LhunzeCounty) BSangngagqoiling三安曲林() feud.Barony {
	return c.三安曲林Sangngagqoiling
}
    
func (c *隆子LhunzeCounty) BYumai玉麦() feud.Barony {
	return c.玉麦Yumai
}
    
func (c *隆子LhunzeCounty) BZhunba准巴() feud.Barony {
	return c.准巴Zhunba
}
    
var CLhunze隆子 LhunzeCounty = &隆子LhunzeCounty{}

func init() {
	f := CLhunze隆子.(*隆子LhunzeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1552",
		Title:     "lhunze",
		TitleName: "隆子",
		TitleCode: "c_lhunze",
		Baronies:  map[string]feud.Barony{},
	}

	f.错那Cona_b = BCona_b错那
	f.错那Cona_b.SetParent(f)
	
	f.隆子Lhunze = BLhunze隆子
	f.隆子Lhunze.SetParent(f)
	
	f.拿日雍Nariyong = BNariyong拿日雍
	f.拿日雍Nariyong.SetParent(f)
	
	f.日当Ritang = BRitang日当
	f.日当Ritang.SetParent(f)
	
	f.三安曲林Sangngagqoiling = BSangngagqoiling三安曲林
	f.三安曲林Sangngagqoiling.SetParent(f)
	
	f.玉麦Yumai = BYumai玉麦
	f.玉麦Yumai.SetParent(f)
	
	f.准巴Zhunba = BZhunba准巴
	f.准巴Zhunba.SetParent(f)
	
}
