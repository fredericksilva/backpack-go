## Pacotes
- [package.go](/package.go)

Cada programa Go � composto de pacotes.
Programas come�am rodando pelo pacote �main�.
Este programa est� usando os pacotes com caminhos de importa��o "fmt" e "math".
Por conve��o, o nome do pacote � o mesmo que o �ltimo elemento do caminho de importa��o.

**Nota**: o ambiente em que esses programas s�o executados � **determin�stico**, ent�o �rad.Intn� sempre retornar� o mesmo n�mero.(Para ver um n�mero diferente, a semente do gerador de n�meros; veja �rand.Seed�.)
