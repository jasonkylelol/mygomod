package main

import (
	"fmt"
	"strings"
)

func genInferRefreshImgName(imgName string, iter int32) string {
	dotList := strings.Split(imgName, ".")
	if len(dotList) <= 1 {
		return fmt.Sprintf("%v.iter%v", imgName, iter)
	}
	postfix := dotList[len(dotList)-1]
	iterinfo := dotList[len(dotList)-2]
	newDotList := []string{}
	if strings.Contains(iterinfo, "iter") {
		dotList[len(dotList)-2] = fmt.Sprintf("iter%v", iter)
		newDotList = append(newDotList, dotList...)
	} else {
		newDotList = append(newDotList, dotList[:len(dotList)-1]...)
		newDotList = append(newDotList, fmt.Sprintf("iter%v", iter), postfix)
	}
	return strings.Join(newDotList, ".")
}

func imgIterReplace() {
	imgNames := []string{
		"aaaaaa.jpg",
		"aaaaaa.iter1.jpg",
		"aaaaaa.iter1.iter1.jpg",
		"P830014151900121112300042868_1211125_1535.res__UM02-1350000050140__yi_wei__Pseudo_Error__735x630__1.jpg",
		"P830014151900121112300042868_1211125_1535.res__UM02-1350000050140__yi_wei__Pseudo_Error__735x630__1.iter1.jpg",
	}
	for idx, name := range imgNames {
		fmt.Println(idx, ":", genInferRefreshImgName(name, 2))
	}
}
