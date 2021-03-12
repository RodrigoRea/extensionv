# extensionv - Golang
Write values ​​in full. Description of a numeric value

Descreve um numero por extenso no tipo moeda brasileira (Real R$)

Limitado ao número de caracteres enviado para a função - Máximo da descrição: trilhões

A solução é retornar um valor do tipo float para a descrição em reais 
Ex: 2.47 => dois reais quarenta e sete centavos

# Install 
<pre>
  go get github.com/RodrigoRea/extensionv
</pre>

# Exemplo de uso 
<pre>
  import "github.com/RodrigoRea/extensionv"
  .
  .
  .
  
  var n float64 = 7850.89  
  vtext, _ := extensionv.Descreva(n)
  
  // Valor descrito contido em `vtext`
	
</pre>
