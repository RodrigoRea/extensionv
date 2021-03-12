# extensionv - Golang
Write values ​​in full. Description of a numeric value

Descreve um numero por extenso no tipo moeda brasileira (Real R$)

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
