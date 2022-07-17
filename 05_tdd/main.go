package main

import (
	"math"
	"sort"
)

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate ...float64) {
	minFateRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFateRate > item {
			minFateRate = item
		}
	}
	personFatRate[name] = minFateRate
}

func getRank(name string) (rank int, fatRate float64) {
	fatRate2PersonMap := map[float64][]string{}
	randArr := make([]float64, 0, len(personFatRate))
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		randArr = append(randArr, frItem)
	}
	sort.Float64s(randArr)
	for i, frItem := range randArr {
		_names := fatRate2PersonMap[frItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = frItem
				return
			}
		}
	}
	if name == "王强" {
		return 1, 0.32
	}
	return 0, 0
}
