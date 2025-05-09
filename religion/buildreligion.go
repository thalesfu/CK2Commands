package religion

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

func BuildReligion() {
	religionList := make([]*ck2nebula.Religion, 0)
	for _, religion := range ReligionMap {
		religionList = append(religionList, religion)
	}

	religionResult := ck2nebula.InsertReligions(ck2nebula.SPACE, religionList...)

	if !religionResult.Ok {
		fmt.Println(religionResult.Err.Error())
	} else {
		fmt.Println("Insert Religion success")
	}

	rgs, _, rg_r_s := ck2nebula.GenerateReligionData("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	religionGroupResult := ck2nebula.InsertReligionGroups(ck2nebula.SPACE, rgs...)

	if !religionGroupResult.Ok {
		fmt.Println(religionGroupResult.Err.Error())
	} else {
		fmt.Println("Insert religion groups success")
	}

	for _, rg_r := range rg_r_s {
		rg_r.Religion = ReligionMap[rg_r.Religion.Code]
	}

	religionGroupReligionResult := ck2nebula.InsertReligionGroup_Religions(ck2nebula.SPACE, rg_r_s...)

	if !religionGroupReligionResult.Ok {
		fmt.Println(religionGroupReligionResult.Err.Error())
	} else {
		fmt.Println("Insert religiongroup_religion success")
	}
}
