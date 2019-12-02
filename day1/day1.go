package main

import (
	"go.uber.org/zap"
)

func main() {
	total := 0
	logger, _ := zap.NewDevelopment()

	for _, mass := range input {
		result := calculatFuel(mass, logger.Sugar())
		logger.Sugar().Infow("total fuel for module", "mass", mass, "fuel", result)

		total += result
	}

	logger.Sugar().Infow("total calculated", "total", total)
}

func calculatFuel(mass int, logger *zap.SugaredLogger) int {
	result := (int(mass/3)) - 2

	if result > 0 {
		result += calculatFuel(result, logger)
		return result
	} else {
		return 0
	}
}


var input = []int {
	72713,
	93795,
	64596,
	99366,
	124304,
	122702,
	105674,
	94104,
	144795,
	109412,
	138753,
	71738,
	62172,
	149671,
	88232,
	145707,
	82617,
	123357,
	63980,
	149016,
	130921,
	125863,
	119405,
	77839,
	140354,
	135213,
	130550,
	131981,
	149301,
	68884,
	52555,
	121036,
	75237,
	64339,
	60612,
	132912,
	63164,
	145198,
	109252,
	130024,
	100738,
	74890,
	89784,
	134474,
	68815,
	117283,
	144774,
	138017,
	149989,
	111506,
	119737,
	65330,
	52388,
	69698,
	124990,
	84232,
	58016,
	142321,
	119731,
	86914,
	68524,
	87708,
	60776,
	119259,
	73429,
	79486,
	120369,
	57007,
	91514,
	87226,
	131770,
	78170,
	52871,
	149060,
	73804,
	60034,
	72519,
	98065,
	132526,
	99660,
	74854,
	111912,
	51232,
	71499,
	127629,
	83807,
	91061,
	60988,
	133493,
	95170,
	110661,
	54486,
	63241,
	141111,
	142244,
	93120,
	137257,
	79822,
	95849,
	69878,
}