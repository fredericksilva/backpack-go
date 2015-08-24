## Table of Contents

- [Importações](#imports)
- [Pacotes](#Package) 

## Imports (Importações)
* File: [Imports.go](/basic/imports.go)

Este grupo de códigos em parênteses da importação, é a declaração de importação **"consignada"**. Você também pode escrever várias declarações de importação assim:

```import "fmt"
   import "math"```

Mas é um ***"design pattern"*** usar a declaração de importação consignada. 

## Package (Pacotes)

* File: [package.go](/basic/packages.go)

Cada programa Go é composto de pacotes.
Programas começam rodando pelo pacote `main`.
Este programa está usando os pacotes com caminhos de importação "fmt" e "math".
Por conveção, o nome do pacote é o mesmo que o último elemento do caminho de importação.

**Nota**: o ambiente em que esses programas são executados é **determinístico**, então `rad.Intn` sempre retornará o mesmo número.(Para ver um número diferente, a semente do gerador de números; veja `rand.Seed`.)

---
:v: @frederickSilva
