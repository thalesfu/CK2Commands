package shigatse

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/gyantse"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/gyesar"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/lhatse"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/mangyul"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/sakya"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse/shigatse"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShigatseDuke interface {
	feud.Duke
	CGyantse江孜() gyantse.GyantseCounty
	CGyesar杰萨() gyesar.GyesarCounty
	CLhatse拉孜() lhatse.LhatseCounty
	CMangyul芒隅() mangyul.MangyulCounty
	CSakya萨迦() sakya.SakyaCounty
	CShigatse日喀则() shigatse.ShigatseCounty
}

type 日喀则ShigatseDuke struct {
	feud.BaseDuke
	江孜Gyantse   gyantse.GyantseCounty
	杰萨Gyesar    gyesar.GyesarCounty
	拉孜Lhatse    lhatse.LhatseCounty
	芒隅Mangyul   mangyul.MangyulCounty
	萨迦Sakya     sakya.SakyaCounty
	日喀则Shigatse shigatse.ShigatseCounty
}

func (d *日喀则ShigatseDuke) CGyantse江孜() gyantse.GyantseCounty {
	return d.江孜Gyantse
}

func (d *日喀则ShigatseDuke) CGyesar杰萨() gyesar.GyesarCounty {
	return d.杰萨Gyesar
}

func (d *日喀则ShigatseDuke) CLhatse拉孜() lhatse.LhatseCounty {
	return d.拉孜Lhatse
}

func (d *日喀则ShigatseDuke) CMangyul芒隅() mangyul.MangyulCounty {
	return d.芒隅Mangyul
}

func (d *日喀则ShigatseDuke) CSakya萨迦() sakya.SakyaCounty {
	return d.萨迦Sakya
}

func (d *日喀则ShigatseDuke) CShigatse日喀则() shigatse.ShigatseCounty {
	return d.日喀则Shigatse
}

var DShigatse日喀则 ShigatseDuke = &日喀则ShigatseDuke{}

func init() {
	f := DShigatse日喀则.(*日喀则ShigatseDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "shigatse",
		TitleName: "日喀则",
		TitleCode: "d_shigatse",
		Counties:  map[string]feud.County{},
	}

	f.江孜Gyantse = gyantse.CGyantse江孜
	f.江孜Gyantse.SetParent(f)

	f.杰萨Gyesar = gyesar.CGyesar杰萨
	f.杰萨Gyesar.SetParent(f)

	f.拉孜Lhatse = lhatse.CLhatse拉孜
	f.拉孜Lhatse.SetParent(f)

	f.芒隅Mangyul = mangyul.CMangyul芒隅
	f.芒隅Mangyul.SetParent(f)

	f.萨迦Sakya = sakya.CSakya萨迦
	f.萨迦Sakya.SetParent(f)

	f.日喀则Shigatse = shigatse.CShigatse日喀则
	f.日喀则Shigatse.SetParent(f)

}
