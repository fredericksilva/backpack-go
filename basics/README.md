## Table of Contents

- [Nomes Exportados](#exported-names)
- [Importações](#imports)
- [Pacotes](#Package) 

## Exported Names (Nomes Exportados)
- File: [exported-names.go](/basics/exported-names.go)

Depois de importar um pacote você pode consultar os nomes que exporta.
Em **Go*, um nome é exportado se começa coma letra maiúscula.
`Foo`é um nome exportado, assim como `FOO`. O nome `foo`não é exportado.


## Imports (Importações)
* File: [Imports.go](/basics/imports.go)

Este grupo de códigos em parênteses da importação, é a declaração de importação **"consignada"**. Você também pode escrever várias declarações de importação assim:

```import "fmt"
   import "math"```

Mas é um ***"design pattern"*** usar a declaração de importação consignada. 

## Package (Pacotes)

* File: [package.go](/basics/packages.go)

Cada programa Go é composto de pacotes.
Programas começam rodando pelo pacote `main`.
Este programa está usando os pacotes com caminhos de importação "fmt" e "math".
Por conveção, o nome do pacote é o mesmo que o último elemento do caminho de importação.

**Nota**: o ambiente em que esses programas são executados é **determinístico**, então `rad.Intn` sempre retornará o mesmo número.(Para ver um número diferente, a semente do gerador de números; veja `rand.Seed`.)

---
:v: @frederickSilva
