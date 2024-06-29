## Hello
Go exige que você importe o package `main` e também o `fmt`. Para que você possa compilar código.

## Importações
Para importar pacotes usando Go é possível desta forma:

```
import {
	"package1"
	"package2"
}
```

Desde modo é possível manter um organização do que foi importado.

## Nome Exportados
Quando um pacote é exportado ele começa com letra maiúcula.

Por exemplo, `math.pi` não é exportado mas `math.Pi` é exportado.

> Todos os nomes "não exportados" não são acessíveis de fora do pacote.

## Funções
Em Golang as funções precisam ser declaradas fora da `main()` e então chamadas.

Para declarar o tipo das entrada Go segue a seguinte lógica.

```
func add(a int b int) int {
	return a + b
}
```

## Links Úteis
- [Go Sintaxe](https://go.dev/blog/declaration-syntax)
- [Go Operations](https://www.geeksforgeeks.org/go-operators/)