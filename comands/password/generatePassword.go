package password

import (
	"math/rand"

	"github.com/pedro-makoski/fast-tools/utils"
)

var charsMaisculo = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var charsMinusculo = []rune("abcdefghijklmnopqrstuvwxyz")
var simbols = []rune("@!?()-_#*")
var numbers = []rune("0123456789")

func (p *PasswordGen) generatePassword() (string, error) {
	senha := []rune{}

	skeleton := p.randomSkeleton()

	for _, v := range skeleton {
		senha = append(senha, randomChar(v))
	}

	return string(senha), nil
}

func randomChar(typeOfChar int) rune {
	switch typeOfChar {
	case MAISCULA:
		return utils.GetRandomValueOfList(charsMaisculo)
	case MINUSCULA:
		return utils.GetRandomValueOfList(charsMinusculo)
	case NUMBER:
		return utils.GetRandomValueOfList(numbers)
	case SIMBOL:
		return utils.GetRandomValueOfList(simbols)
	}

	return '*'
}

func (p *PasswordGen) randomSkeleton() []int {
	values := []int{}

	lengtOfEach := p.genTotalToUseForEach()

	for typeOfLength, quant := range lengtOfEach {
		for range quant {
			values = append(values, typeOfLength)
		}
	}

	return utils.Permutate(values)
}

func (p *PasswordGen) genTotalToUseForEach() map[int]int {
	var totalToUse = map[int]int{
		MAISCULA:  0,
		MINUSCULA: 0,
		NUMBER:    0,
		SIMBOL:    0,
	}

	totalUsed := 0
	quantTested := 0

	totalItems := utils.BoolToInt(p.hasNumbers) +
		utils.BoolToInt(p.hasSimbols) +
		(utils.BoolToInt(p.hasChars) * utils.BoolToInt(p.isMaiusculo)) +
		(utils.BoolToInt(p.hasChars) * utils.BoolToInt(p.isMinusculo))

	if p.isMaiusculo && p.hasChars {
		totalToUse[MAISCULA] = rand.Intn(p.length-(totalItems-quantTested)) + 1
		totalUsed += totalToUse[MAISCULA]
		quantTested++
	}

	if p.isMinusculo && p.hasChars {
		totalToUse[MINUSCULA] = rand.Intn((p.length-totalUsed)-(totalItems-quantTested)) + 1
		totalUsed += totalToUse[MINUSCULA]
		quantTested++
	}

	if p.hasSimbols {
		totalToUse[SIMBOL] = rand.Intn((p.length-totalUsed)-(totalItems-quantTested)) + 1
		totalUsed += totalToUse[SIMBOL]
		quantTested++
	}

	if p.hasNumbers {
		totalToUse[NUMBER] = p.length - totalUsed
	}

	return totalToUse
}
