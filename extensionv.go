package extensionv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var unidade = []string{
	"",
	"um",
	"dois",
	"três",
	"quatro",
	"cinco",
	"seis",
	"sete",
	"oito",
	"nove",
}

var dezena10 = []string{
	"dez",
	"onze",
	"doze",
	"treze",
	"quatorze",
	"quinze",
	"dezesseis",
	"dezesete",
	"dezoito",
	"dezenove",
}

var dezena = []string{
	"",
	"dez",
	"vinte",
	"trinta",
	"quarenta",
	"cinquenta",
	"sessenta",
	"setenta",
	"oitenta",
	"noventa",
}

var centena = []string{
	"",
	"cem",
	"duzentos",
	"trezentos",
	"quatrocentos",
	"quinhentos",
	"seiscentos",
	"setecentos",
	"oitocentos",
	"novecentos",
}

func descritivo(sNumber string) (texto string) {

	n, err := strconv.Atoi(sNumber)
	if err != nil {
		return
	}
	if n <= 0 {
		return
	}

	switch {
	case n < 10:
		texto = unidade[n]
	case n < 20:
		texto = dezena10[n-10]

	case n < 100:
		sn1 := sNumber[len(sNumber)-2 : len(sNumber)-1]
		sn2 := sNumber[len(sNumber)-1 : len(sNumber)]
		n1, _ := strconv.Atoi(sn1)
		n2, _ := strconv.Atoi(sn2)
		texto = dezena[n1]
		if n2 > 0 {
			texto += " e " + unidade[n2]
		}
	case n == 100:
		texto = centena[1]
	case n > 100 && n < 200:
		sn2 := sNumber[len(sNumber)-2 : len(sNumber)-1]
		sn3 := sNumber[len(sNumber)-1 : len(sNumber)]

		n2, _ := strconv.Atoi(sn2)
		n3, _ := strconv.Atoi(sn3)
		texto = "cento"
		if n2 > 0 {
			d, _ := strconv.Atoi(sn2 + sn3)
			if d > 9 && d < 20 {
				texto += " e " + descritivo(sn2+sn3)
				n3 = 0
			} else if d < 100 {
				texto += " e " + descritivo(sn2+sn3)
				n3 = 0
			} else {
				texto += " e " + descritivo(sn2)
			}
		}
		if n3 > 0 {
			texto += " e " + descritivo(sn3)
		}

	case n < 1000:
		sn1 := sNumber[len(sNumber)-3 : len(sNumber)-2]
		sn2 := sNumber[len(sNumber)-2 : len(sNumber)-1]
		sn3 := sNumber[len(sNumber)-1 : len(sNumber)]

		n1, _ := strconv.Atoi(sn1)
		n2, _ := strconv.Atoi(sn2)
		n3, _ := strconv.Atoi(sn3)
		texto = centena[n1]
		if n2 > 0 {
			d, _ := strconv.Atoi(sn2 + sn3)
			if d > 9 && d < 20 {
				texto += " e " + descritivo(sn2+sn3)
				n3 = 0
			} else if d < 100 {
				texto += " e " + descritivo(sn2+sn3)
				n3 = 0
			} else {
				texto += " e " + descritivo(sn2)
			}
		}
		if n3 > 0 {
			texto += " e " + descritivo(sn3)
		}

	case n < 1000000:
		sn1 := sNumber[0 : len(sNumber)-3]
		sn2 := sNumber[len(sNumber)-3 : len(sNumber)]
		texto += descritivo(sn1) + " mil"
		n2, _ := strconv.Atoi(sn2)
		if n2 > 0 {
			texto += " e " + descritivo(sn2)
		}

	case n < 1000000000:
		sn1 := sNumber[0 : len(sNumber)-6]
		desc := " milhão"
		n1, _ := strconv.Atoi(sn1)
		if n1 > 1 {
			desc = " milhões"
		}
		texto += descritivo(sn1) + desc
		sn2 := sNumber[len(sNumber)-6 : len(sNumber)]
		n2, _ := strconv.Atoi(sn2)
		if n2 > 0 {
			texto += " e " + descritivo(sn2)
		}
	
	case n < 1000000000000:
		sn1 := sNumber[0 : len(sNumber)-9]
		desc := " bilhão"
		n1, _ := strconv.Atoi(sn1)
		if n1 > 1 {
			desc = " bilhões"
		}
		texto += descritivo(sn1) + desc
		sn2 := sNumber[len(sNumber)-9 : len(sNumber)]
		n2, _ := strconv.Atoi(sn2)
		if n2 > 0 {
			texto += " e " + descritivo(sn2)
		}
	case n < 1000000000000000:
		sn1 := sNumber[0 : len(sNumber)-12]
		desc := " trilhão"
		n1, _ := strconv.Atoi(sn1)
		if n1 > 1 {
			desc = " trilhões"
		}
		texto += descritivo(sn1) + desc
		sn2 := sNumber[len(sNumber)-12 : len(sNumber)]
		n2, _ := strconv.Atoi(sn2)
		if n2 > 0 {
			texto += " e " + descritivo(sn2)
		}
	case n < 1000000000000000000:
		sn1 := sNumber[0 : len(sNumber)-15]
		desc := " quintilhão"
		n1, _ := strconv.Atoi(sn1)
		if n1 > 1 {
			desc = " quintilhões"
		}
		texto += descritivo(sn1) + desc
		sn2 := sNumber[len(sNumber)-12 : len(sNumber)]
		n2, _ := strconv.Atoi(sn2)
		if n2 > 0 {
			texto += " e " + descritivo(sn2)
		}
	}
	return

}

// Numerico - Retorna a descrição do valor por extenso em Numerico
func Numerico(n int) (texto string) {

	v := float64(n)
	extenso, _ := Descreva(v)
	extenso = strings.Replace(extenso, "reais", "", -1)
	extenso = strings.Replace(extenso, "real", "", -1)
	extenso = strings.Replace(extenso, "centavos", "", -1)
	extenso = strings.Replace(extenso, "centavo", "", -1)

	extenso = strings.Trim(extenso, "")
	return extenso
}

// Descreva - Retorna a descrição do valor por extenso em Real
func Descreva(n float64) (texto string, err error) {

	// Converter para string
	n, err = strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	if err != nil {
		return
	}
	sNumber := strconv.FormatFloat(n, 'f', 2, 64)

	// Deixar apenas números
	regx := regexp.MustCompile(`[^0-9]`)
	sNumber = regx.ReplaceAllString(sNumber, "")

	centavos := sNumber[len(sNumber)-2 : len(sNumber)]
	valor := sNumber[0 : len(sNumber)-2]

	nCentavos, _ := strconv.Atoi(centavos)
	nValor, _ := strconv.Atoi(valor)

	centavos = descritivo(centavos)
	valor = descritivo(valor)

	if nValor == 0 && nCentavos == 0 {
		return "zero real", nil
	}

	r := ""
	if nValor > 0 {
		r = " reais"
		if nValor == 1 {
			r = " real"
		}
	}

	c := ""
	if nCentavos > 0 {
		if nValor > 0 {
			c += " e " + centavos
		} else {
			c += centavos
		}
		if nCentavos == 1 {
			c += " centavo"
		} else {
			c += " centavos"
		}
	}

	sNumber = valor + r + c
	sNumber = strings.Replace(sNumber, "mil e", "mil", -1)
	sNumber = strings.Replace(sNumber, "milhão e", "milhão", -1)
	sNumber = strings.Replace(sNumber, "milhões e", "milhões", -1)
	sNumber = strings.Replace(sNumber, "bilhões e", "bilhões", -1)
	sNumber = strings.Replace(sNumber, "trilhões e", "trilhões", -1)
	sNumber = strings.Replace(sNumber, "quintilhões e", "quintilhões", -1)
	// fmt.Printf("\n Descritivo do numero %+v \n %s \n", fmt.Sprintf("%.2f", n), sNumber)
	// return "", errors.New("Err")
	return sNumber, nil

}
