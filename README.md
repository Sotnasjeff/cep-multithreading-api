# Exercicio

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/cep

http://viacep.com.br/ws/" + cep + "/json/

# Requisitos do Exercicio

Os requisitos para este desafio são:

- [x] Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- [x] O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- [x] Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

# Requisitos Adicionais (criado pelo próprio dev)
- [x] Realizar testes unitários.

# Requisitos para rodar a aplicação e como rodar?
- Possuir o Golang na versão 1.20

- Para rodar a aplicação somente rodar "go run main.go" na pasta cmd/server

- Uma vez que o Server estiver rodando, você pode ir na pasta "test" e testar a aplicação.