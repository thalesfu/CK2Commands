package ikh_bogd

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/ikh_bogd/burkhan_buudai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/ikh_bogd/delgerkhangai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/ikh_bogd/gurvan_saikhan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/ikh_bogd/ikh_bogd"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Ikh_bogdDuke interface {
    feud.Duke
    CBurkhan_buudai布尔汗布岱() 	burkhan_buudai.Burkhan_buudaiCounty
    CDelgerkhangai德勒格尔杭爱() 	delgerkhangai.DelgerkhangaiCounty
    CGurvan_saikhan古尔班赛罕() 	gurvan_saikhan.Gurvan_saikhanCounty
    CIkh_bogd也客孛黑答() 	ikh_bogd.Ikh_bogdCounty
}

type 也客孛黑答Ikh_bogdDuke struct {
	feud.BaseDuke
	布尔汗布岱Burkhan_buudai 	burkhan_buudai.Burkhan_buudaiCounty
	德勒格尔杭爱Delgerkhangai 	delgerkhangai.DelgerkhangaiCounty
	古尔班赛罕Gurvan_saikhan 	gurvan_saikhan.Gurvan_saikhanCounty
	也客孛黑答Ikh_bogd 	ikh_bogd.Ikh_bogdCounty
}

func (d *也客孛黑答Ikh_bogdDuke) CBurkhan_buudai布尔汗布岱() burkhan_buudai.Burkhan_buudaiCounty {
	return d.布尔汗布岱Burkhan_buudai
}
    
func (d *也客孛黑答Ikh_bogdDuke) CDelgerkhangai德勒格尔杭爱() delgerkhangai.DelgerkhangaiCounty {
	return d.德勒格尔杭爱Delgerkhangai
}
    
func (d *也客孛黑答Ikh_bogdDuke) CGurvan_saikhan古尔班赛罕() gurvan_saikhan.Gurvan_saikhanCounty {
	return d.古尔班赛罕Gurvan_saikhan
}
    
func (d *也客孛黑答Ikh_bogdDuke) CIkh_bogd也客孛黑答() ikh_bogd.Ikh_bogdCounty {
	return d.也客孛黑答Ikh_bogd
}
    
var DIkh_bogd也客孛黑答 Ikh_bogdDuke = &也客孛黑答Ikh_bogdDuke{}

func init() {
	f := DIkh_bogd也客孛黑答.(*也客孛黑答Ikh_bogdDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ikh_bogd",
		TitleName: "也客孛黑答",
		TitleCode: "d_ikh_bogd",
		Counties:  map[string]feud.County{},
	}

	f.布尔汗布岱Burkhan_buudai = burkhan_buudai.CBurkhan_buudai布尔汗布岱
	f.布尔汗布岱Burkhan_buudai.SetParent(f)
	
	f.德勒格尔杭爱Delgerkhangai = delgerkhangai.CDelgerkhangai德勒格尔杭爱
	f.德勒格尔杭爱Delgerkhangai.SetParent(f)
	
	f.古尔班赛罕Gurvan_saikhan = gurvan_saikhan.CGurvan_saikhan古尔班赛罕
	f.古尔班赛罕Gurvan_saikhan.SetParent(f)
	
	f.也客孛黑答Ikh_bogd = ikh_bogd.CIkh_bogd也客孛黑答
	f.也客孛黑答Ikh_bogd.SetParent(f)
	
}
