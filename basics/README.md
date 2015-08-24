## Table of Contents

- [Nomes Exportados](#exported-names)
- [Importa��es](#imports)
- [Pacotes](#Package) 

## Exported Names (Nomes Exportados)
- File: [exported-names.go](/basics/exported-names.go)

Depois de importar um pacote voc� pode consultar os nomes que exporta.
Em **Go*, um nome � exportado se come�a coma letra mai�scula.
`Foo`� um nome exportado, assim como `FOO`. O nome `foo`n�o � exportado.


## Imports (Importa��es)
* File: [Imports.go](/basics/imports.go)

Este grupo de c�digos em par�nteses da importa��o, � a declara��o de importa��o **"consignada"**. Voc� tamb�m pode escrever v�rias declara��es de importa��o assim:

```import "fmt"
   import "math"```

Mas � um ***"design pattern"*** usar a declara��o de importa��o consignada. 

## Package (Pacotes)

* File: [package.go](/basics/packages.go)

Cada programa Go � composto de pacotes.
Programas come�am rodando pelo pacote `main`.
Este programa est� usando os pacotes com caminhos de importa��o "fmt" e "math".
Por conve��o, o nome do pacote � o mesmo que o �ltimo elemento do caminho de importa��o.

**Nota**: o ambiente em que esses programas s�o executados � **determin�stico**, ent�o `rad.Intn` sempre retornar� o mesmo n�mero.(Para ver um n�mero diferente, a semente do gerador de n�meros; veja `rand.Seed`.)

---
:v: @frederickSilva
