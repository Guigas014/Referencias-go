## Jogo da forca

- primeiro criaamos um modulo do go, uando o seguinte comando:

      `go mod init github.com/Guigas014/forca`

- depois é necessário roda o comando:

      `go mod tidy`

Obs.: o arquivo "go.mod" foi criado.

<br />

### Echo framework

Vamos utilizar o framework ECHO para criar um web service.
Para instalar ele use o seguinte comando:

      `go get github.com/labstack/echo/v4`

E depois rode o "go mod tidy" novamente.

Obs.: o arquivo "go.sum" foi criado.

<br />

### Passos para criar o jogo

1. Na função "main" chamamos a função "initGame()" para iniciar o jogo e depois criamos três rotas: um get para mostrar a palavra, um post para processar a letra enviada pela request e outro get para reiniciar o jogo.

2. Cada requisição chama uma função:

   - GET - /word -> getWord()
   - POST - /word -> insertWord()
   - GET - /word/new -> newGame()

3. Para guardar as palavras e o estado delas, criamos uma "struct" chamada GameState.

4. Para guardar a letra que vem da request POST foi criado a "struct" Request. Funciona como o BodyRequest do Java.

5. Foi criado também uma variável chamda gameState, a qual é do tipo GameState e é por ela por ela que gerenciamos as palavras.

6. A função "getWord()" apenas retorna a palavra guardada no estado **showWord**.

7. A função "insertWord()" captura a letra enviada pela requisição e se não acontecer nenhum a função "processLetter()" é chamada.

8. A função "processLetter()" recebe a letra, passa para minúsculo caso ela esteja maiúscula, testa se a letra existe na palavra. Se sim substitui a letra no estado **showWord** que é retornado para o usuário, se não testa se a palavra já foi toda descoberta. Se a palavra já foi toda descoberta, um texto de parabéns e a revelação da palavra é retornada ao usuário.

9. A função "newGame()" limpa o estado **showWord** e inicia o jogo novamente chamando a função "initGame()".

10. A função "initGame()" por sua vez popula o slice **words** com as palavras a serem usadas no jogo. Depois gera um índice aleatório para sortear a palavra da vez. Seta o estado **chosenWord** com a palavra sorteada e popula o estado **showWord** com um underline para cada letra.
